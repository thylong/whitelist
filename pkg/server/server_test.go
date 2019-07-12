package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	router := NewServer()
	router.GET("/healthz", Healthz)

	r, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.FailNow()
	}
}

func TestInsertIP(t *testing.T) {
	router := NewServer()
	router.POST("/ip", InsertIP)

	r, _ := http.NewRequest("POST", "/ip", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.FailNow()
	}
}

func TestDeleteIP(t *testing.T) {
	router := NewServer()
	router.DELETE("/ip", DeleteIP)
	router.POST("/ip", InsertIP)

	rInsert, _ := http.NewRequest("POST", "/ip", nil)
	rDelete, _ := http.NewRequest("DELETE", "/ip", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()

	// Delete on empty whitelist should return a 500 error
	router.ServeHTTP(w1, rDelete)
	if w1.Code != http.StatusInternalServerError {
		t.FailNow()
	}

	// Insert an IP in the whitelist
	router.ServeHTTP(w2, rInsert)

	// Delete an existing ip should return a 200
	router.ServeHTTP(w3, rDelete)
	if w3.Code != http.StatusOK {
		t.FailNow()
	}
}

func TestContainIP(t *testing.T) {
	router := NewServer()
	router.GET("/ip", ContainIP)
	router.POST("/ip", InsertIP)

	r1, _ := http.NewRequest("GET", "/ip", nil)
	r2, _ := http.NewRequest("POST", "/ip", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()

	// The whitelist is empty at this point
	router.ServeHTTP(w1, r1)
	if w1.Code != http.StatusNotFound {
		t.FailNow()
	}

	// Insert an IP in the whitelist
	router.ServeHTTP(w2, r2)

	router.ServeHTTP(w3, r1)
	if w3.Code != http.StatusOK {
		t.FailNow()
	}
}

func sendInsertRequest(client *http.Client, addr string) {
	res, err := client.Post(addr, "application/json", nil)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		panic("request failed")
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
}
