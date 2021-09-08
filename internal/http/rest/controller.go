package rest

import (
	"net/http"

	"github.com/sss-eda/lemi-011b/internal/lemi011b"
)

// Lemi011bController TODO
type Lemi011bController struct {
	service *lemi011b.Service
	mux     *http.ServeMux
}

// NewLemi011bController TODO
func NewLemi011bController(
	service *lemi011b.Service,
) (*Lemi011bController, error) {
	mux := http.NewServeMux()

	mux.Handle("/device", &DeviceHandler{})
	mux.Handle("/devices", &DevicesHandler{})

	return &Lemi011bController{
		service: service,
		mux:     mux,
	}, nil
}

// ServeHTTP TODO
func (ctrl *Lemi011bController) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

}
