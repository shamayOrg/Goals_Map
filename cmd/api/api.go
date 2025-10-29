package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	port string
}

func (app *application) mountRouts() *http.ServeMux {
	mux := http.NewServeMux()
	// Route handlers would be registered here
	mux.HandleFunc("/v1/health", app.healthCheckHandler)
	return mux
}

func (app *application) start(mux *http.ServeMux) error {
	// Implementation to start the API server would go here
	srv := &http.Server{
		Addr:         ":" + app.config.port,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started in port %s", app.config.port)

	return srv.ListenAndServe()
}
