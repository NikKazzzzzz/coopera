package main

import (
	"log"

	"github.com/NikKazzzzzz/coopera-backend/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
}
