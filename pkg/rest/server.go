package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/registration"
)

// API TODO
type API struct {
	acquirer acquisition.Service
}

// AcquireDatum TODO
func AcquireDatumHandler(
	acquirer acquisition.Service,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println(fmt.Errorf("can only POST to this endpoint"))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		datum := acquisition.Datum{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&datum)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = acquirer.AcquireDatum(r.Context(), datum)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// RegisterSensorHandler TODO
func RegisterSensorHandler(
	registry registration.Service,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println(fmt.Errorf("can only POST to this endpoint"))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		sensor := registration.Sensor{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&sensor)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = registry.RegisterSensor(r.Context(), sensor)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
