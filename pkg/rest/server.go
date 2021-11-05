package rest

import (
	"context"
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

// AcquireDatumHandler TODO
func AcquireDatumHandler(
	callback func(context.Context, acquisition.Datum) error,
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

		err = callback(r.Context(), datum)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// RegisterInstrumentHandler TODO
func RegisterInstrumentHandler(
	callback func(context.Context, registration.Instrument) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println(fmt.Errorf("can only POST to this endpoint"))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		instrument := registration.Instrument{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&instrument)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = callback(r.Context(), instrument)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
