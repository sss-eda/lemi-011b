package main

import (
	"log"

	"github.com/sss-eda/lemi-011b/internal/lemi011b"
	"github.com/sss-eda/lemi-011b/internal/physical"
	"github.com/tarm/serial"
)

func main() {
	device1, err := serial.OpenPort(
		&serial.Config{
			Name: "/dev/ttyUSB0",
			Baud: 115200,
		},
	)
	if err != nil {
		log.Fatalf("unable to open serial port: %v", err)
	}
	defer device1.Close()

	repository, err := physical.NewDeviceRepository(device1)
	if err != nil {
		log.Fatalf("unable to create repository: %v", err)
	}

	presenter, err := local.NewDatumPresenter()

	service, err := lemi011b.NewService(repository, presenter)
	if err != nil {
		log.Fatalf("unable to create repository: %v", err)
	}

	log.Fatal(local.Log(service))
}
