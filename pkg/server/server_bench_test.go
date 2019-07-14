// +build benchmark

package server

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startDummyServer() {
	if httpServer != nil {
		return
	}

	httpServer = &http.Server{
		Handler: &testHandler{},
	}

	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}

	go func() {
		err := httpServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
}

func BenchmarkFilterWithRadixSingleIPWhitelist(b *testing.B) {
	startDummyServer()

	backendHost = "localhost:8083"
	defer func() { backendHost = "dummy:8080" }()

	router := NewServer("radix")
	router.GET("/", Filter)
	router.POST("/ip/:ip", InsertIP)

	// Seed the trie to prepare deletion tests
	rInsert, _ := http.NewRequest("POST", "/ip/203.0.113.195", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)

	r, _ := http.NewRequest("GET", "http://localhost:8083/", nil)
	r.Header.Add("X-Forwarded-For", "203.0.113.195")

	client := &http.Client{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res, err := client.Do(r)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		if res.StatusCode != 200 {
			b.Fatalf("wrong status code: %d\n", res.StatusCode)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		err = res.Body.Close()
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}
	}
}

func BenchmarkFilterWithRadixLargeIPWhitelist(b *testing.B) {
	startDummyServer()

	backendHost = "localhost:8083"
	defer func() { backendHost = "dummy:8080" }()

	router := NewServer("radix")
	router.GET("/", Filter)
	router.POST("/ip/:ip", InsertIP)

	go http.ListenAndServe(":8080", router)

	// Seed the trie to prepare for large search accross the structure
	for i := 0; i < 100000; i++ {
		space1 := rand.Intn(255)
		space2 := rand.Intn(255)
		space3 := rand.Intn(255)
		space4 := rand.Intn(255)

		ip := fmt.Sprintf("%d%d%d%d", space1, space2, space3, space4)

		rInsert, _ := http.NewRequest("POST", fmt.Sprintf("/ip/%s", ip), nil)
		wInsert := httptest.NewRecorder()
		router.ServeHTTP(wInsert, rInsert)
	}

	// Whitelist IP for the test (to not go through 403)
	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	rInsert, _ := http.NewRequest("POST", "/ip/203.0.113.195", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)
	r.Header.Add("X-Forwarded-For", "203.0.113.195")

	client := &http.Client{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res, err := client.Do(r)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		if res.StatusCode != 200 {
			b.Fatalf("StatusCode: %d\n", res.StatusCode)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		err = res.Body.Close()
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}
	}
}

func BenchmarkFilterWithListSingleIPWhitelist(b *testing.B) {
	startDummyServer()

	backendHost = "localhost:8083"
	defer func() { backendHost = "dummy:8080" }()

	router := NewServer("list")
	router.GET("/", Filter)
	router.POST("/ip/:ip", InsertIP)

	// Seed the trie to prepare deletion tests
	rInsert, _ := http.NewRequest("POST", "/ip/203.0.113.195", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)

	r, _ := http.NewRequest("GET", "http://localhost:8083/", nil)
	r.Header.Add("X-Forwarded-For", "203.0.113.195")

	client := &http.Client{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res, err := client.Do(r)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		if res.StatusCode != 200 {
			b.Fatalf("wrong status code: %d\n", res.StatusCode)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		err = res.Body.Close()
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}
	}
}

func BenchmarkFilterWithListLargeIPWhitelist(b *testing.B) {
	startDummyServer()

	backendHost = "localhost:8083"
	defer func() { backendHost = "dummy:8080" }()

	router := NewServer("list")
	router.GET("/", Filter)
	router.POST("/ip/:ip", InsertIP)

	go http.ListenAndServe(":8080", router)

	// Seed the trie to prepare for large search accross the structure
	for i := 0; i < 100000; i++ {
		space1 := rand.Intn(255)
		space2 := rand.Intn(255)
		space3 := rand.Intn(255)
		space4 := rand.Intn(255)

		ip := fmt.Sprintf("%d%d%d%d", space1, space2, space3, space4)

		rInsert, _ := http.NewRequest("POST", fmt.Sprintf("/ip/%s", ip), nil)
		wInsert := httptest.NewRecorder()
		router.ServeHTTP(wInsert, rInsert)
	}

	// Whitelist IP for the test (to not go through 403)
	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	rInsert, _ := http.NewRequest("POST", "/ip/203.0.113.195", nil)
	wInsert := httptest.NewRecorder()
	router.ServeHTTP(wInsert, rInsert)
	r.Header.Add("X-Forwarded-For", "203.0.113.195")

	client := &http.Client{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res, err := client.Do(r)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		if res.StatusCode != 200 {
			b.Fatalf("StatusCode: %d\n", res.StatusCode)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}

		err = res.Body.Close()
		if err != nil {
			b.Fatalf("error: %s\n", err)
		}
	}
}
