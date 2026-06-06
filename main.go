package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Psyduck000054/chirpy/functions"
	"github.com/Psyduck000054/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Can't open database")
		return
	}
	dbQueries := database.New(db)

	apiCfg := functions.ApiConfig{
		FileServerHits: atomic.Int32{},
		Queries:        dbQueries,
		Platform:       platform,
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

	mux.HandleFunc("POST /api/chirps", apiCfg.HandlerCreateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.HandlerCreateUser)
	mux.HandleFunc("GET /api/healthz", functions.HandlerReadiness)
	mux.HandleFunc("GET /api/chirps", apiCfg.HandlerRetrieveChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.HandlerRetrieveChirp)

	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerReset)

	// starts listening
	server.ListenAndServe()
}
