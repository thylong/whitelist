package server

import (
	"net/http"

	"github.com/thylong/whitelist/pkg/whitelist"

	"github.com/julienschmidt/httprouter"
)

var tree *whitelist.Tree

// PanicHandler prevent server crash on panic serving 500 error instead
func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.WriteHeader(500)
}

// Healthz returns 200 to health checks
func Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

// InsertIP adds a given IP to the whitelist
func InsertIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := tree.Insert(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(500)
	}
}

// DeleteIP deletes a given IP of the whitelist
func DeleteIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := tree.Delete(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(500)
	}
}

// ContainIP returns 200 if a given IP is in the whitelist
func ContainIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := tree.Contain(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(404)
	}
}

// Filter out not expected requests and forward to the next service
func Filter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

// NewServer starts a new HTTP server
func NewServer() *httprouter.Router {
	tree = whitelist.New()
	return httprouter.New()
}
