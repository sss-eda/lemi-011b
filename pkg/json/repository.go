package json

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/configuration"
)

type repository struct{}

// NewRepository TODO
func NewRepository() (configuration.Repository, error) {
	return &repository{}, nil
}

// Configure TODO
func (repo *repository) Configure(
	ctx context.Context,
	config interface{},
) error {
	return nil
}
