package main

import (
	"context"
	"flag"
	"log"

	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/rest"
	"github.com/sss-eda/lemi-011b/pkg/serial"

	tarm "github.com/tarm/serial"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := flag.String("api_url", "", "URL of API server.")
	id := flag.Uint64("instrument_id", 0, "Instrument ID.")
	serialName := flag.String("serial_name", "/dev/ttyUSB0", "Serial port name.")
	serialBaud := flag.Int("serial_baude", 19200, "Serial port baud rate.")

	flag.Parse()

	client, err := rest.NewClient(*url)
	if err != nil {
		log.Fatal(err)
	}

	acquirer, err := acquisition.NewService(client)
	if err != nil {
		log.Fatal(err)
	}

	controller, err := serial.NewController(*id, acquirer)
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: *serialName,
		Baud: *serialBaud,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(controller.Run(ctx, port))
}
