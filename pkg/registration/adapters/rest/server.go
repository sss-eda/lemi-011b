package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/registration"
)

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
