package main

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/prototype/routes/game"
)

func SetupRoutes() {
	http.HandleFunc("/game/socket/", game.SocketHandler)

	http.ListenAndServe(":8080", nil)
}

func main() {
	SetupRoutes()
}
