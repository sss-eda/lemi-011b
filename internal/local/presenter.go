package local

import (
	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// Presenter TODO
type Presenter struct{}

// NewPresenter TODO
func NewPresenter() (*Presenter, error) {
	return &Presenter{}, nil
}

// Present TODO
func (pres *Presenter) Present(
	datum *lemi011b.Datum,
) error {
	// Save to file
	return nil
}

// Log TODO
func Log(service lemi011b.Service) error {
	return nil
}
