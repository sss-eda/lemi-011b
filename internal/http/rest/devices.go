package rest

import (
	"fmt"
	"net/http"

	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// DevicesService TODO
type DevicesService interface {
	ListDevices() ([]*lemi011b.Device, error)
}

// DevicesHandler TODO
type DevicesHandler struct {
	service DevicesService
}

// ServeHTTP TODO
func (handler *DevicesHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		handler.get(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed,
			fmt.Errorf("endpoint only supports the GET method"))
	}
}

func (handler *DevicesHandler) get(
	w http.ResponseWriter,
	r *http.Request,
) {
	devices, err := handler.service.ListDevices()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	}

}
