package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var tr *http.Transport
var client *http.Client
var httpServer *http.Server

type testHandler struct{}

func (th *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK.\n"))
}

func TestHealthz(t *testing.T) {
	router := NewServer("radix")
	router.GET("/healthz", Healthz)

	r, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.FailNow()
	}
}

func TestInsertIP(t *testing.T) {
	router := NewServer("radix")
	router.POST("/ip/:ip", InsertIP)

	tr = &http.Transport{
		MaxIdleConns:          50,
		IdleConnTimeout:       2 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
	}
	client = &http.Client{Transport: tr}

	testCases := []struct {
		Method         string
		Path           string
		ExpectedStatus int
	}{
		{Method: "POST", Path: "/ip/127.0.0.1", ExpectedStatus: 200},      // Test successful insertion
		{Method: "POST", Path: "/ip/127.0.0.1", ExpectedStatus: 500},      // Test duplicate insertion
		{Method: "POST", Path: "/ip/127.0.0.133333", ExpectedStatus: 400}, // Test invalid client insertion
	}
	for _, tc := range testCases {
		r, _ := http.NewRequest(tc.Method, tc.Path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		if w.Code != tc.ExpectedStatus {
			t.FailNow()
		}
	}
}

func TestDeleteIP(t *testing.T) {
	router := NewServer("radix")
	router.DELETE("/ip/:ip", DeleteIP)
	router.POST("/ip/:ip", InsertIP)

	tr = &http.Transport{
		MaxIdleConns:          50,
		IdleConnTimeout:       2 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
	}
	client = &http.Client{Transport: tr}

	// Seed the trie to prepare deletion tests
	rInsert, _ := http.NewRequest("POST", "/ip/127.0.0.1", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)

	testCases := []struct {
		Method         string
		Path           string
		ExpectedStatus int
	}{
		{Method: "DELETE", Path: "/ip/127.0.0.1", ExpectedStatus: 200},      // Test successful deletion
		{Method: "DELETE", Path: "/ip/127.0.0.1", ExpectedStatus: 500},      // Test duplicate deletion
		{Method: "DELETE", Path: "/ip/127.0.0.133333", ExpectedStatus: 400}, // Test invalid client deletion
	}
	for _, tc := range testCases {
		r, _ := http.NewRequest(tc.Method, tc.Path, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		if w.Code != tc.ExpectedStatus {
			t.FailNow()
		}
	}
}

func TestContainIP(t *testing.T) {
	router := NewServer("radix")
	router.GET("/ip/:ip", ContainIP)
	router.POST("/ip/:ip", InsertIP)

	// Seed the trie to prepare deletion tests
	rInsert, _ := http.NewRequest("POST", "/ip/127.0.0.1", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)

	testCases := []struct {
		Method         string
		Path           string
		ExpectedStatus int
	}{
		{Method: "GET", Path: "/ip/127.0.0.1", ExpectedStatus: 200},      // Test successful lookup
		{Method: "GET", Path: "/ip/127.0.0.20", ExpectedStatus: 404},     // Test failed lookup
		{Method: "GET", Path: "/ip/127.0.0.133333", ExpectedStatus: 400}, // Test invalid client lookup
	}
	for _, tc := range testCases {
		r, _ := http.NewRequest(tc.Method, tc.Path, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		if w.Code != tc.ExpectedStatus {
			t.FailNow()
		}
	}
}

func TestFilter(t *testing.T) {
	router := NewServer("radix")
	router.GET("/", Filter)
	router.POST("/ip/:ip", InsertIP)

	testCases := []struct {
		Method         string
		Allow          bool
		Path           string
		ExpectedStatus int
	}{
		{Method: "GET", Allow: false, Path: "http://localhost/", ExpectedStatus: 403}, // Test denied proxy
		{Method: "GET", Allow: true, Path: "http://localhost/", ExpectedStatus: 200},  // Test allowed proxy
	}
	for _, tc := range testCases {
		backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "this call was relayed by the reverse proxy")
		}))
		defer backendServer.Close()

		rpURL, err := url.Parse(backendServer.URL)
		if err != nil {
			log.Fatal(err)
		}

		backendHost = rpURL.Host
		defer func() { backendHost = "dummy:8080" }()

		r, _ := http.NewRequest(tc.Method, tc.Path, nil)
		if tc.Allow {
			// process IP address
			rInsert, _ := http.NewRequest("POST", "/ip/203.0.113.195", nil)
			wInsert := httptest.NewRecorder()
			router.ServeHTTP(wInsert, rInsert)

			r.Header.Add("X-Forwarded-For", "203.0.113.195")
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		if w.Code != tc.ExpectedStatus {
			t.FailNow()
		}
	}
}
