package main

import (
	"goals_map/internal/env"
	"log"
)

func main() {
	cfg := config{
		port: env.GetEnv("PORT", "8080"),
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
