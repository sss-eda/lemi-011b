package configuration

import "context"

// Service TODO
type Service interface {
	Configure(context.Context, interface{}) error
}

// Repository TODO
type Repository interface {
	Configure(context.Context, interface{}) error
}
