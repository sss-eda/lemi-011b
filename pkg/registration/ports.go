package registration

import "context"

// Service TODO
type Service interface {
	RegisterInstrument(context.Context, Instrument) error
}

// Repository TODO
type Repository interface {
	RegisterInstrument(context.Context, Instrument) error
}
