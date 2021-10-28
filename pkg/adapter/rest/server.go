package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/domain/registration"
)

// Server TODO
type Server struct {
	mux      *http.ServeMux
	acquirer acquisition.Service
	registry registration.Service
}

// NewServer TODO
func NewServer(
	acquisitionService acquisition.Service,
	registrationService registration.Service,
) (*Server, error) {
	server := &Server{
		mux:      http.NewServeMux(),
		acquirer: acquisitionService,
		registry: registrationService,
	}

	// server.mux.HandleFunc("/", http.HandlerFunc(http.FileServer(http.Dir(".\\web")).ServeHTTP))
	server.mux.Handle("/", http.FileServer(http.Dir("./web")))
	server.mux.HandleFunc("/datum", server.AcquireDatum)
	server.mux.HandleFunc("/sensor", server.RegisterSensor)

	return server, nil
}

// ServeHTTP TODO
func (server *Server) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	server.mux.ServeHTTP(w, r)
}

// AcquireDatum TODO
func (server *Server) AcquireDatum(
	w http.ResponseWriter,
	r *http.Request,
) {
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

	err = server.acquirer.AcquireDatum(r.Context(), datum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RegisterSensor TODO
func (server *Server) RegisterSensor(
	w http.ResponseWriter,
	r *http.Request,
) {
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

	err = server.registry.RegisterSensor(r.Context(), sensor)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
