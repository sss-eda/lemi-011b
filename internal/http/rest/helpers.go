package rest

import (
	"log"
	"net/http"
)

func respondWithError(
	w http.ResponseWriter,
	status int,
	err error,
) {
	w.WriteHeader(status)
	n, err := w.Write([]byte("{\"error: " + err.Error() + "}"))
	log.Printf("failed to respond with error: %v, wrote %d bytes", err, n)
}
