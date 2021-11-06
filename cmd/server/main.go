package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/http"
	"github.com/sss-eda/lemi-011b/pkg/https"
	"github.com/sss-eda/lemi-011b/pkg/registration"
	"github.com/sss-eda/lemi-011b/pkg/rest"
	"github.com/sss-eda/lemi-011b/pkg/timescale"
)

// flag value
type tlsHostsVar []string

// String TODO
func (hosts *tlsHostsVar) String() string {
	return "my string representation"
}

// Set TODO
func (hosts *tlsHostsVar) Set(value string) error {
	*hosts = append(*hosts, value)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpPort := flag.Uint64("port", 8080, "Port to serve HTTP on. Exclusive of TLS.")
	tlsEnabled := flag.Bool("tls", false, "Enable TLS.")
	tlsCertDir := flag.String("tls_dir", "", "TLS certificate cache directory.")

	tlsHosts := tlsHostsVar{}
	flag.Var(&tlsHosts, "tls_host", "TLS allowed host name.")

	flag.Parse()

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

	handler, err := rest.NewHandler(acquirer, registrar)
	if err != nil {
		log.Fatal(err)
	}

	if *tlsEnabled {
		log.Fatal(https.Serve(https.Config{
			Hosts:   tlsHosts,
			CertDir: *tlsCertDir,
		}, handler))
	} else {
		log.Fatal(http.Serve(http.Config{
			Port: *httpPort,
		}, handler))
	}
}
