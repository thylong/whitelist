package server

import (
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/thylong/whitelist/pkg/ip"
	"github.com/thylong/whitelist/pkg/whitelist"

	"github.com/julienschmidt/httprouter"
)

var storage whitelist.Storage

var backendHost = "dummy:80"

// PanicHandler prevent server crash on panic serving 500 error instead
func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.WriteHeader(500)
}

// Healthz returns 200 to health checks
func Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

// InsertIP adds a given IP to the whitelist
func InsertIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Filter out invalid IPv4 & IPv6 IPs
	ipAddress := net.ParseIP(ps.ByName("ip"))
	if ipAddress == nil {
		w.WriteHeader(400)
		return
	}

	ok := storage.Insert(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(500)
	}
}

// DeleteIP deletes a given IP of the whitelist
func DeleteIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Filter out invalid IPv4 & IPv6 IPs
	ipAddress := net.ParseIP(ps.ByName("ip"))
	if ipAddress == nil {
		w.WriteHeader(400)
		return
	}

	ok := storage.Delete(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(500)
	}
}

// ContainIP returns 200 if a given IP is in the whitelist
func ContainIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Filter out invalid IPv4 & IPv6 IPs
	ipAddress := net.ParseIP(ps.ByName("ip"))
	if ipAddress == nil {
		w.WriteHeader(400)
		return
	}

	ok := storage.Contain(ps.ByName("ip"))
	if !ok {
		w.WriteHeader(404)
	}
}

// Filter out not expected requests and forward to the next service
func Filter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ip := ip.FindIP(r)
	ok := storage.Contain(ip)
	if !ok {
		w.WriteHeader(403)
	}

	// TODO: Make dynamic and configurable
	r.URL.Scheme = "http"
	r.URL.Host = backendHost
	r.Host = backendHost

	// create the reverse proxy & serve request to backend
	proxy := httputil.NewSingleHostReverseProxy(r.URL)

	proxy.ServeHTTP(w, r)
}

// NewServer starts a new HTTP server
func NewServer(storageKind string) *httprouter.Router {
	storage = whitelist.New(storageKind)
	return httprouter.New()
}
