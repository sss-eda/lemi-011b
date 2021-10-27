package registration

import "context"

// Repository TODO
type Repository interface {
	RegisterSensor(context.Context, Sensor) error
}

// Service TODO
type Service interface {
	RegisterSensor(context.Context, Sensor) error
}
