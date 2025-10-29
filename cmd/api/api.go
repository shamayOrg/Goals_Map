package main

import (
	"net/http"
)

type application struct {
	config config
}

type config struct {
	port string
}

func (app *application) start() error {
	// Implementation to start the API server would go here
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    app.config.port,
		Handler: mux,
	}
	return srv.ListenAndServe()
}
