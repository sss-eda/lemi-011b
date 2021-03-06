package serial

import (
	"bufio"
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	tarm "github.com/tarm/serial"
)

// Controller TODO
type Controller struct {
	instrumentID acquisition.InstrumentID
	service      acquisition.Service
}

// NewController TODO
func NewController(
	instrumentID acquisition.InstrumentID,
	acquisitionService acquisition.Service,
) (*Controller, error) {
	return &Controller{
		instrumentID: instrumentID,
		service:      acquisitionService,
	}, nil
}

// Run TODO
func (ctrl *Controller) Run(
	ctx context.Context,
	port *tarm.Port,
) error {
	scanner := bufio.NewScanner(port)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ", ")
		if len(fields) != 4 {
			continue
		}
		timestamp := time.Now()
		x, err := strconv.ParseInt(fields[0], 10, 32)
		if err != nil {
			log.Println(err)
			continue
		}
		y, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			log.Println(err)
			continue
		}
		z, err := strconv.ParseInt(fields[2], 10, 32)
		if err != nil {
			log.Println(err)
			continue
		}
		t, err := strconv.ParseInt(fields[3], 10, 16)
		if err != nil {
			log.Println(err)
			continue
		}

		datum := acquisition.Datum{
			Time:         timestamp,
			InstrumentID: ctrl.instrumentID,
			X:            int32(x),
			Y:            int32(y),
			Z:            int32(z),
			T:            int16(t),
		}

		err = ctrl.service.AcquireDatum(ctx, datum)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return scanner.Err()
}
