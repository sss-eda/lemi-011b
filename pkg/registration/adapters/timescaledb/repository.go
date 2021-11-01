package timescaledb

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/registration"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository TODO
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository TODO
func NewRepository(
	ctx context.Context,
	pgxpool *pgxpool.Pool,
) (*Repository, error) {
	repo := &Repository{
		pool: pgxpool,
	}

	_, err := repo.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS sensor
		(
			id SERIAL PRIMARY KEY
		);

		CREATE TABLE IF NOT EXISTS datum
		(
			time TIMESTAMPTZ PRIMARY KEY,
			sensor_id INTEGER,
			x INTEGER,
			y INTEGER,
			z INTEGER,
			t INTEGER,
			FOREIGN KEY (sensor_id) REFERENCES sensor (id)
		);

		SELECT create_hypertable('datum', 'time', if_not_exists => TRUE);
	`)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// RegisterSensor TODO
func (repo *Repository) RegisterSensor(
	ctx context.Context,
	sensor registration.Sensor,
) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO sensor (id)
		VALUES ($1);
	`, sensor.ID)
	if err != nil {
		return err
	}

	return nil
}
