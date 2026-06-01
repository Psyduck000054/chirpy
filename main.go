package main

import (
	"net/http"
	"sync/atomic"

	"github.com/Psyduck000054/chirpy/functions"
)

func main() {
	apiCfg := functions.ApiConfig{
		FileServerHits: atomic.Int32{},
	}

	// create a multiplexer
	mux := http.NewServeMux()

	// configure the server
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// set it up
	handler := http.FileServer(http.Dir("./app"))
	strippedHandler := http.StripPrefix("/app/", handler)

	// registers the route
	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(strippedHandler))

	mux.HandleFunc("GET /api/healthz", functions.HandlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerReset)

	// starts listening
	server.ListenAndServe()
}
