package main

import (
	"log"
)

func main() {
	cfg := config{
		port: "8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mountRoutes()
	err := app.start(mux)
	if err != nil {
		log.Fatal(err)
	}
}
