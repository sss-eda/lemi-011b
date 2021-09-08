package lemi011b

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

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
		line := scanner.Text()

		fields := strings.Split(line, ", ")
		timestamp := time.Now()
		// TODO: Error handling
		x, _ := strconv.ParseInt(fields[0], 10, 64)
		y, _ := strconv.ParseInt(fields[1], 10, 64)
		z, _ := strconv.ParseInt(fields[2], 10, 64)
		t, _ := strconv.ParseInt(fields[3], 10, 64)

		datum := Datum{
			Timestamp: timestamp,
			X:         x,
			Y:         y,
			Z:         z,
			T:         t,
		}

		err = svc.data.Present(&datum)
		if err != nil {
			log.Printf("failed to present datum: %v", datum)
		}
	}

	return fmt.Errorf("unexpected error: %v", scanner.Err())
}
