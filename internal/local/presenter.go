package local

import (
	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// DatumPresenter TODO
type DatumPresenter struct{}

// NewDatumPresenter TODO
func NewDatumPresenter() (*DatumPresenter, error) {
	return &DatumPresenter{}, nil
}

// Present TODO
func (pres *DatumPresenter) Present(
	datum *lemi011b.Datum,
) error {
	// Save to file
	return nil
}
