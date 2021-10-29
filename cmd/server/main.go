package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sss-eda/lemi-011b/pkg/adapter/rest"
	"github.com/sss-eda/lemi-011b/pkg/adapter/timescaledb"
	"github.com/sss-eda/lemi-011b/pkg/core"
)

func main() {
	ctx := context.Background()

	timescaledbURL := os.Getenv("LEMI011B_SERVER_TIMESCALEDB_URL")
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

	acquirer, err := core.NewAcquisitionService(repo)
	if err != nil {
		log.Fatal(err)
	}

	registry, err := core.NewRegistrationService(repo)
	if err != nil {
		log.Fatal(err)
	}

	server, err := rest.NewServer(acquirer, registry)
	if err != nil {
		log.Fatal(err)
	}

	parseFlags()
	var m *autocert.Manager

	var httpsSrv *http.Server
	if flgProduction {
		hostPolicy := func(ctx context.Context, host string) error {
			// Note: change to your real host
			allowedHost := "sansa.dev"
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
		}

		dataDir := "."
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(dataDir),
		}

		httpsSrv = makeServerFromMux(server)
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
	if flgRedirectHTTPToHTTPS {
		httpSrv = makeHTTPToHTTPSRedirectServer()
	} else {
		httpSrv = makeServerFromMux(server)
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

// func main() {

// 	// TODO: If there aren't any certs
// 	//  -> Generate some self-signed ones?
// 	//  -> Just use HTTP instead?
// 	//  -> Panic?
// 	log.Fatal(http.ListenAndServeTLS(":443", "/certs/fullchain.pem", "/certs/privkey.pem", server))
// }
