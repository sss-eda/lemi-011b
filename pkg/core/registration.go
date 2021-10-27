package core

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/domain/registration"
)

// RegistrationService TODO
type RegistrationService struct {
	repo registration.Repository
}

// NewRegistrationService TODO
func NewRegistrationService(
	repository registration.Repository,
) (*RegistrationService, error) {
	return &RegistrationService{
		repo: repository,
	}, nil
}

// RegisterSensor TODO
func (service *RegistrationService) RegisterSensor(
	ctx context.Context,
	sensor registration.Sensor,
) error {
	return service.repo.RegisterSensor(ctx, sensor)
}
