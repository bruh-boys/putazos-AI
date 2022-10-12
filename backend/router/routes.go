package router

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/backend/controller"
)

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	http.HandleFunc("/game", controller.SeeGame)
	http.HandleFunc("/Multiplayer-connection", controller.Multiplayer)

	http.ListenAndServe(":8080", nil)

}
