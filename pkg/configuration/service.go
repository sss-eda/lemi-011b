package configuration

import "context"

type service struct{}

// NewService TODO
func NewService(
	ctx context.Context,
) (Service, error) {
	return &service{}, nil
}

// ParseENV TODO
func (svc *service) ParseENV(
	ctx context.Context,
	config interface{},
) error {
	return nil
}
