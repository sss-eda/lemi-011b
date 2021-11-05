package acquisition

import "context"

// Service TODO
type Service struct {
	repo Repository
}

// NewService TODO
func NewService(
	repository Repository,
) (*Service, error) {
	return &Service{
		repo: repository,
	}, nil
}

// AcquireDatum TODO
func (service *Service) AcquireDatum(
	ctx context.Context,
	datum Datum,
) error {
	return service.repo.AcquireDatum(ctx, datum)
}
