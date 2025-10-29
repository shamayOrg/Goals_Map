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

	err := app.start()
	if err != nil {
		log.Fatal(err)
	}
}
