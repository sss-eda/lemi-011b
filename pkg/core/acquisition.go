package core

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
)

// AcquisitionService TODO
type AcquisitionService struct {
	repo acquisition.Repository
}

// NewAcquisitionService TODO
func NewAcquisitionService(
	repository acquisition.Repository,
) (*AcquisitionService, error) {
	return &AcquisitionService{
		repo: repository,
	}, nil
}

// AcquireDatum TODO
func (service *AcquisitionService) AcquireDatum(
	ctx context.Context,
	datum acquisition.Datum,
) error {
	return service.repo.AcquireDatum(ctx, datum)
}
