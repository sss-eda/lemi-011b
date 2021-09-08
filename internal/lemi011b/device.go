package lemi011b

import (
	"io"

	"github.com/google/uuid"
)

// DeviceID TODO
type DeviceID uuid.UUID

// Device TODO
type Device struct {
	ID     DeviceID
	Reader io.Reader
}
