package main

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/sss-eda/lemi-011b/internal/lemi011b"
	"github.com/sss-eda/lemi-011b/internal/local"
	"github.com/sss-eda/lemi-011b/internal/physical"
	"github.com/tarm/serial"
)

func main() {
	uuid1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("failed to generate ID for device1")
	}
	id1 := lemi011b.DeviceID(uuid1)
	port1, err := serial.OpenPort(
		&serial.Config{
			Name: "/dev/ttyUSB0",
			Baud: 115200,
		},
	)
	if err != nil {
		log.Fatalf("unable to open serial port: %v", err)
	}
	defer port1.Close()

	device1 := &lemi011b.Device{
		ID:     lemi011b.DeviceID(id1),
		Reader: port1,
	}

	repository, err := physical.NewDeviceRepository(device1)
	if err != nil {
		log.Fatalf("unable to create repository: %v", err)
	}

	presenter, err := local.NewDatumPresenter()

	service, err := lemi011b.NewService(repository, presenter)
	if err != nil {
		log.Fatalf("unable to create repository: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func(id lemi011b.DeviceID, wg *sync.WaitGroup) {
		log.Fatal(service.AcquireData(id))
	}(id1, &wg)

	wg.Wait()
}
