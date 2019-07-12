package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fooHandler := func(w http.ResponseWriter, req *http.Request) {
		var reqHeaders bytes.Buffer
		for k, v := range req.Header {
			reqHeaders.WriteString(fmt.Sprint(k, ":", v, "\n"))
		}
		fmt.Println(reqHeaders.String())
		io.WriteString(w, "foo like you!\n")
	}
	barHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "bar like you!\n")
	}
	foobarHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "foobar like you!\n")
	}
	healthcheckHandler := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	}

	http.HandleFunc("/", fooHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/foobar", foobarHandler)
	http.HandleFunc("/healthcheck", healthcheckHandler)
	http.HandleFunc("/healthcheck-pastrami", healthcheckHandler)
	http.HandleFunc("/healthz", healthcheckHandler)

	fmt.Println("Listening on :80")
	http.ListenAndServe(":80", nil)
}
