package lemi011b

import "bufio"

// Service TODO
type Service struct {
	devices DeviceRepository
	data    DatumPresenter
}

// NewService TODO
func NewService(
	deviceRepository DeviceRepository,
	datumPresenter DatumPresenter,
) (*Service, error) {
	return &Service{
		devices: deviceRepository,
		data:    datumPresenter,
	}, nil
}

// AcquireData TODO
func (svc *Service) AcquireData(
	id DeviceID,
) error {
	device, err := svc.devices.Load(id)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(device.Reader)
	for scanner.Scan() {
		s := scanner.Text()
		svc.data.Present(s)
	}
}
