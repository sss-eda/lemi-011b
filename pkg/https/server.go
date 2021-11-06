package https

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

// Server TODO
func Serve(
	config Config,
	handler http.Handler,
) error {
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(config.Hosts...),
		Cache:      autocert.DirCache(config.CertDir),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	fmt.Printf("Starting HTTPS server on %s\n", server.Addr)
	go server.ListenAndServeTLS("", "")

	return http.ListenAndServe(":80", certManager.HTTPHandler(nil))
}
