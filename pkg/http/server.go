package http

import (
	"fmt"
	"net/http"
	"strconv"
)

// Serve TODO
func Serve(
	config Config,
	handler http.Handler,
) error {
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(config.Port), 10),
		Handler: handler,
	}

	fmt.Printf("Starting HTTP server on port:%v\n", config.Port)
	return server.ListenAndServe()
}
