package registration

import "context"

// Service TODO
type Service interface {
	RegisterSensor(context.Context, Sensor) error
}

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

// RegisterSensor TODO
func (svc *service) RegisterSensor(
	ctx context.Context,
	sensor Sensor,
) error {
	return svc.repo.RegisterSensor(ctx, sensor)
}
