package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang/glog"
)

/**
 * Url examples:
 * 1. http://localhost:9000
 * 2. http://localhost:9000/?user=james
 * 3. http://localhost:9000/healthz
 */
func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server")

	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/healthz", healthz)

	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns
	// and calls the handler for the pattern that most closely matches the URL.
	// In electronics, a multiplexer (or mux; spelled sometimes as multiplexor), also known as a data selector,
	// is a device that selects between several analog or digital input signals and forwards the selected input to a single output line.
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)

	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)

	// Listens on port 9000
	// err := http.ListenAndServe(":9000", nil)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Root url handler, namely "/"
 */
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===============Details of the http request header:===============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

/**
 * HealthZ handler, namely "/healthz"
 */
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}
