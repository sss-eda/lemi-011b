package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/sss-eda/lemi-011b/pkg/adapter/rest"
	"github.com/sss-eda/lemi-011b/pkg/adapter/serial"
	"github.com/sss-eda/lemi-011b/pkg/core"
	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"

	tarm "github.com/tarm/serial"
)

func main() {
	ctx := context.Background()

	envSensorID := os.Getenv("LEMI011B_CLIENT_SENSOR_ID")
	if envSensorID == "" {
		log.Fatal("No environment variable for sensor ID")
	}
	sensorID, err := strconv.ParseInt(envSensorID, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	restURL := os.Getenv("LEMI011B_CLIENT_REST_URL")
	if restURL == "" {
		log.Fatal("No environment variable for rest url")
	}
	serialName := os.Getenv("LEMI011B_CLIENT_SERIAL_PORT")
	if serialName == "" {
		log.Fatal("No environment variable for serial port name")
	}
	serialBaud := os.Getenv("LEMI011B_CLIENT_SERIAL_BAUD")
	if serialBaud == "" {
		log.Fatal("No environment variable for serial baud")
	}
	serialBaudInt, err := strconv.Atoi(serialBaud)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := rest.NewClient(restURL)
	if err != nil {
		log.Fatal(err)
	}

	service, err := core.NewAcquisitionService(repo)
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: serialName,
		Baud: serialBaudInt,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctrl, err := serial.NewController(acquisition.SensorID(sensorID), service)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(ctrl.Run(ctx, port))
}
