package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var tr *http.Transport
var client *http.Client

func main() {
	tr = &http.Transport{
		MaxIdleConns:          50,
		IdleConnTimeout:       2 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
	}
	client = &http.Client{Transport: tr}

	testCases := []struct {
		Method         string
		BaseURI        string
		Path           string
		ExpectedStatus int
	}{
		{Method: "GET", BaseURI: "http://whitelist:8080", Path: "/", ExpectedStatus: 403}, // Denied
		{Method: "POST", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 200},
		// {Method: "GET", BaseURI: "http://whitelist:8080", Path: "/", ExpectedStatus: 200}, // Allowed
		{Method: "POST", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 500},
		{Method: "GET", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 200},
		{Method: "DELETE", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 200},
		{Method: "DELETE", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 500},
		{Method: "GET", BaseURI: "http://whitelist:8081", Path: "/ip/127.0.0.1", ExpectedStatus: 404},
	}
	for _, tc := range testCases {
		// Format test request
		req, err := http.NewRequest(tc.Method, tc.BaseURI+tc.Path, nil)
		if err != nil {
			log.Fatal(err)
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("HTTP request failed, status: %d,\nexpected: %d\n", res.StatusCode, tc.ExpectedStatus)
		}
		defer res.Body.Close()

		// Check Response status code
		if res.StatusCode != tc.ExpectedStatus {
			log.Fatalf("Test %#v failed, status received %d, expected %d\n", tc, res.StatusCode, tc.ExpectedStatus)
		}
		fmt.Printf("Test %#v successful\n", tc)
	}
}
