package main

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
