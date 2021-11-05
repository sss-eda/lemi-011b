package acquisition

import "context"

// Service TODO
// type Service interface {
// 	AcquireDatum(context.Context, Datum) error
// }

// AcquireDatumUseCase TODO
type AcquireDatumUseCase func(context.Context, Datum) error

// Repository TODO
type Repository interface {
	AcquireDatum(context.Context, Datum) error
}
