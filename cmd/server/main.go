package main

import (
	"context"
	"log"
	"net/http"
	"os"

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
	restPort := os.Getenv("LEMI011B_SERVER_REST_PORT")
	if restPort == "" {
		log.Fatal("no env variable defined for rest port")
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

	log.Fatal(http.ListenAndServe(":"+restPort, server))
}
