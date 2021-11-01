package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/acquisition/adapters/rest"
	"github.com/sss-eda/lemi-011b/pkg/acquisition/adapters/timescaledb"
	"golang.org/x/crypto/acme/autocert"
)

const (
	httpPort = "127.0.0.1:8080"
)

func main() {
	ctx := context.Background()

	timescaledbURL := os.Getenv("TIMESCALEDB_URL")
	if timescaledbURL == "" {
		log.Fatal("no env variable defined for timescaledb url")
	}

	dbpool, err := pgxpool.Connect(ctx, timescaledbURL)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := timescaledb.NewRepository(ctx, dbpool)
	if err != nil {
		log.Fatal(err)
	}

	acquirer, err := acquisition.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	// registry, err := registration.NewService(repo)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server, err := rest.NewServer(acquirer)
	if err != nil {
		log.Fatal(err)
	}

	var m *autocert.Manager

	var httpsSrv *http.Server

	flgProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION"))
	if err != nil {
		log.Fatal(err)
	}
	if flgProduction {
		hostPolicy := func(ctx context.Context, host string) error {
			// Note: change to your real host
			allowedHost := "sansa.dev"
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf(
				"acme/autocert: only %s host is allowed",
				allowedHost,
			)
		}

		dataDir := "."
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(dataDir),
		}

		httpsSrv = &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      server,
		}
		httpsSrv.Addr = ":443"
		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", httpsSrv.Addr)
			err := httpsSrv.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
			}
		}()
	}

	var httpSrv *http.Server

	flgRedirectHTTPToHTTPS, err := strconv.ParseBool(os.Getenv("REDIRECT_TO_HTTPS"))
	if err != nil {
		log.Fatal(err)
	}
	if flgRedirectHTTPToHTTPS {
		httpSrv = mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			newURI := "https://" + r.Host + r.URL.String()
			http.Redirect(w, r, newURI, http.StatusFound)
		})
		return &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      &http.ServeMux{},
		}
	} else {
		httpSrv = &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      server,
		}
	}
	// allow autocert handle Let's Encrypt callbacks over http
	if m != nil {
		httpSrv.Handler = m.HTTPHandler(httpSrv.Handler)
	}

	httpSrv.Addr = httpPort
	fmt.Printf("Starting HTTP server on %s\n", httpPort)
	err = httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
	}
}
