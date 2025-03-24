package main

import (
	"fmt"
	"net/http"
)

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.RequestURI
	http.Redirect(w, r, target, http.StatusMovedPermanently) // 301 Redirect
}

func main() {
	// Start an HTTP server that redirects all requests to HTTPS
	fmt.Println("Starting HTTP redirect server on port 80...")
	httpServer := &http.Server{
		Addr:    ":80",
		Handler: http.HandlerFunc(redirectToHTTPS),
	}

	// Run the HTTP redirect server
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("HTTP server error:", err)
	}
}
