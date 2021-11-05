package timescale

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/registration"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository TODO
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository TODO
func NewRepository(
	pgxPool *pgxpool.Pool,
) (*Repository, error) {
	_, err := pgxPool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS instrument
		(
			id SERIAL PRIMARY KEY
		);

		CREATE TABLE IF NOT EXISTS datum
		(
			time TIMESTAMPTZ PRIMARY KEY,
			instrument_id INTEGER,
			x INTEGER,
			y INTEGER,
			z INTEGER,
			t INTEGER,
			FOREIGN KEY (instrument_id) REFERENCES instrument (id)
		);

		SELECT create_hypertable('datum', 'time', if_not_exists => TRUE);
	`)
	if err != nil {
		return nil, err
	}

	return &Repository{
		pool: pgxPool,
	}, nil
}

// AcquireDatum TODO
func (repo *Repository) AcquireDatum(
	ctx context.Context,
	datum acquisition.Datum,
) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO datum (time, instrument_id, x, y, z, t)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, datum.Time, datum.InstrumentID, datum.X, datum.Y, datum.Z, datum.T)
	if err != nil {
		return err
	}

	return nil
}

// RegisterInstrument TODO
func (repo *Repository) RegisterInstrument(
	ctx context.Context,
	instrument registration.Instrument,
) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO instrument (id)
		VALUES ($1);
	`, instrument.ID)
	if err != nil {
		return err
	}

	return nil
}
