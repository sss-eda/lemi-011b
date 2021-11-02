package configuration

import "context"

// Service TODO
type Service interface {
	ParseENV(context.Context, interface{}) error
	// ParseJSON(context.Context, interface{}) error
	// ParseYAML(context.Context, interface{}) error
}

// Repository TODO
type Repository interface {
	Configure(context.Context, interface{}) error
}
