package router

import (
	"log"
	"net/http"

	"github.com/bruh-boys/putazos-ai/backend/controller"
)

func SetupRoutes(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/game/gateway", controller.Multiplayer)
	http.HandleFunc("/game", controller.SeeGame)

	log.Println("Listening on port " + port)
	return http.ListenAndServe((":" + port), nil)
}
