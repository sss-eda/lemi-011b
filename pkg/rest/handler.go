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

// NewHandler TODO
func NewHandler(
	acquirer acquisition.Service,
	registrar registration.Service,
) (http.Handler, error) {
	mux := http.DefaultServeMux

	fs := http.FileServer(http.Dir("/web"))

	mux.Handle("/", fs)
	mux.HandleFunc("/datum", acquireDatum(acquirer.AcquireDatum))
	mux.HandleFunc("/instrument", registerInstrument(registrar.RegisterInstrument))

	return mux, nil
}

func acquireDatum(
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

func registerInstrument(
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
