package main

import (
	"log"

	"github.com/c4miloarriagada/keys-be/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatalf("Critical error: failed to start the application: %v", err)
	}
}
