package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thylong/whitelist/pkg/server"
)

func main() {
	router := server.NewServer("radix")
	router.PanicHandler = server.PanicHandler

	// Interact with whitelist content
	r := httprouter.New()
	r.POST("/ip/:ip", server.InsertIP)
	r.DELETE("/ip/:ip", server.DeleteIP)
	r.GET("/ip/:ip", server.ContainIP)
	r.GET("/healthz", server.Healthz)

	// Catch all routes
	router.GET("/*path", server.Filter)
	router.PUT("/*path", server.Filter)
	router.DELETE("/*path", server.Filter)
	router.POST("/*path", server.Filter)

	go http.ListenAndServe(":8081", r)
	log.Fatal(http.ListenAndServe(":8080", router))
}
