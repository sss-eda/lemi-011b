package physical

import (
	"fmt"

	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// DeviceRepository TODO
type DeviceRepository struct {
	devices map[lemi011b.DeviceID]*Device
}

// NewDeviceRepository TODO
func NewDeviceRepository(
	devices ...*Device,
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
	device, ok := repo.devices[id]
	if !ok {
		return nil, fmt.Errorf("device with id: %v does not exist", id)
	}

	return &lemi011b.Device{
		ID:     id,
		Reader: device.Reader,
	}, nil
}

// Save TODO
func (repo *DeviceRepository) Save(
	device *lemi011b.DeviceID,
) (lemi011b.DeviceID, error) {
	repo.devices[device.ID] = &Device{
		Reader: device.Reader,
	}
	if !ok {
		return nil, fmt.Errorf("device with id: %v does not exist", id)
	}

	return &lemi011b.Device{
		ID:     id,
		Reader: device.Reader,
	}, nil
}
