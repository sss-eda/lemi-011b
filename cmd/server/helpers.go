package main

import (
	"net/http"
	"time"
)

const (
	httpPort = "127.0.0.1:8080"
)

func makeServerFromMux(mux *http.ServeMux) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPToHTTPSRedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleRedirect)
	return makeServerFromMux(mux)
}

// func parseFlags() {
// 	flag.BoolVar(&flgProduction, "production", false, "if true, we start HTTPS server")
// 	flag.BoolVar(&flgRedirectHTTPToHTTPS, "redirect-to-https", false, "if true, we redirect HTTP to HTTPS")
// 	flag.Parse()
// }
