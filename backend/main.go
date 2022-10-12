package main

import (
	"log"
	"os"

	"github.com/bruh-boys/putazos-ai/backend/router"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"

	}

	log.Panicln(
		router.SetupRoutes(port),
	)

}
