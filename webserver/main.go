package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", serverResponse)
	http.ListenAndServe(":8080", nil)
}

func serverResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", time.Now())
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "HTTP method: %s\n", r.Method)
	user, pass, _ := r.BasicAuth()
	fmt.Fprintf(w, "Basic auth user: %s, \n", user)
	fmt.Fprintf(w, "Basic auth pass: %s, \n", pass)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	// fmt.Fprintf(w, "TLS: %s\n", r.TLS.)
	fmt.Fprintf(w, "UserAgent: %s\n", r.UserAgent())
	fmt.Fprintf(w, "Referer: %s\n", r.Referer())
	fmt.Fprintf(w, "Headers:\n")
	for key, val := range r.Header {
		fmt.Fprintf(w, "  %s=%s\n", key, val)
	}
}
