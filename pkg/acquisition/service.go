package acquisition

import "context"

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

// AcquireDatum TODO
func (svc *service) AcquireDatum(
	ctx context.Context,
	datum Datum,
) error {
	return nil
}
