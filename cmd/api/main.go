package main

import (
	"goals_map/internal/env"

	"go.uber.org/zap"
)

func main() {

	// logger := zap.Must(zap.NewProduction()).Sugar()
	logger := zap.Must(zap.NewDevelopment()).Sugar()
	defer logger.Sync() // flushes buffer, if any

	cfg := config{
		port: env.GetEnv("PORT", "8080"),
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := app.mountRoutes()
	err := app.start(mux)
	if err != nil {
		logger.Fatal(err)
	}
}
