package registration

import (
	"context"
)

// Service TODO
type service struct {
	repo Repository
}

// NewService TODO
func NewService(
	repository Repository,
) (Service, error) {
	return &service{
		repo: repository,
	}, nil
}

// RegisterInstrument TODO
func (svc *service) RegisterInstrument(
	ctx context.Context,
	instrument Instrument,
) error {
	return svc.repo.RegisterInstrument(ctx, instrument)
}
