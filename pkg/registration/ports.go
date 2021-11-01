package registration

import "context"

// Service TODO
type Service interface {
	RegisterSensor(context.Context, Sensor) error
}

// Repository TODO
type Repository interface {
	RegisterSensor(context.Context, Sensor) error
}
