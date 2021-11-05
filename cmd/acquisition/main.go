package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/registration"
	"github.com/sss-eda/lemi-011b/pkg/rest"
	"github.com/sss-eda/lemi-011b/pkg/timescale"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flagSet := flag.NewFlagSet("lemi-011b", flag.ExitOnError)

	tlsEnabled := flagSet.Bool("tls", false, "Enable TLS.")
	tlsCertDir := flagSet.String("tls_dir", "", "TLS certificate cache directory.")

	addr := []string{}
	flagSet.Func("a", "Address to bind to.", func(s string) error {
		if s != "" {
			addr = append(addr, s)
		}
		return fmt.Errorf("no value for --addr flag specified")
	})
	flagSet.Func("addr", "Address to bind to.", func(s string) error {
		if s != "" {
			addr = append(addr, s)
		}
		return fmt.Errorf("no value for --addr flag specified")
	})

	portDefault := uint(8080)
	portUsage := "Port to serve HTTP at (exclusive of TLS)."
	port := flagSet.Uint("port", portDefault, portUsage)
	port := flagSet.Uint("p", portDefault, portUsage)

	timescaleURL := os.Getenv("TIMESCALE_URL")

	dbpool, err := pgxpool.Connect(ctx, timescaleURL)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := timescale.NewRepository(dbpool)
	if err != nil {
		log.Fatal(err)
	}

	acquirer, err := acquisition.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	registrar, err := registration.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	handler := http.DefaultServeMux

	handler.HandleFunc("/datum", rest.AcquireDatumHandler(acquirer.AcquireDatum))
	handler.HandleFunc("/instrument", rest.RegisterInstrumentHandler(registrar.RegisterInstrument))

	if *tlsEnabled {
		certManager := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(addr...),
			Cache:      autocert.DirCache(*tlsCertDir),
		}

		server := &http.Server{
			Addr:    ":443",
			Handler: handler,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", server.Addr)
			err := server.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
			}
		}()
	}

	fmt.Printf("starting HTTP server on %s\n")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
