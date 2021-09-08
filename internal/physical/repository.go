package physical

import (
	"github.com/sss-eda/lemi011b"
)

// DeviceRepository TODO
type DeviceRepository struct {
	devices map[lemi011b.DeviceID]*Device
}

// NewDeviceRepository TODO
func NewDeviceRepository(
	devices ...*lemi011b.Device,
) (*DeviceRepository, error) {
	repo := DeviceRepository{
		devices: map[lemi011b.DeviceID]*Device{},
	}

	return &repo, nil
}

// Load TODO
func (repo *DeviceRepository) Load(
	id lemi011b.DeviceID,
) (*lemi011b.Device, error) {
	device, ok := repo.Load(id)
	if !ok {

	}
}
