package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// Device TODO
type Device struct {
	ID   lemi011b.DeviceID
	Name string
	Baud int
}

// DeviceService TODO
type DeviceService interface {
	AddDevice(*lemi011b.Device) error
}

// DeviceHandler TODO
type DeviceHandler struct {
	service DeviceService
}

// ServeHTTP TODO
func (handler *DeviceHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		handler.get(w, r)
	case http.MethodPost:
		handler.post(w, r)
	}
}

func (handler *DeviceHandler) get(
	w http.ResponseWriter,
	r *http.Request,
) {
	return
}

func (handler *DeviceHandler) post(
	w http.ResponseWriter,
	r *http.Request,
) {
	decoder := json.NewDecoder(r.Body)
	device := &Device{}
	decoder.Decode(device)

	err := handler.service.AddDevice(&lemi011b.Device{
		ID: device.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}
}
