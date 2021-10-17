package serial

import (
	"bufio"
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
	tarm "github.com/tarm/serial"
)

// Controller TODO
type Controller struct {
	sensorID uint64
	service  acquisition.Service
}

// NewController TODO
func NewController(
	sensorID uint64,
	acquisitionService acquisition.Service,
) (*Controller, error) {
	return &Controller{
		sensorID: sensorID,
		service:  acquisitionService,
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
		timestamp := time.Now()
		x, err := strconv.ParseInt(fields[0], 10, 32)
		if err != nil {
			return err
		}
		y, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			return err
		}
		z, err := strconv.ParseInt(fields[2], 10, 32)
		if err != nil {
			return err
		}
		t, err := strconv.ParseInt(fields[3], 10, 16)
		if err != nil {
			return err
		}

		ctrl.service.AcquireDatum(ctx, acquisition.Datum{
			Time:     timestamp,
			SensorID: ctrl.sensorID,
			X:        int32(x),
			Y:        int32(y),
			Z:        int32(z),
			T:        int16(t),
		})
	}

	return nil
}
