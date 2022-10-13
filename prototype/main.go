package main

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/prototype/routes/game"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	router := mux.NewRouter()

	router.HandleFunc(`/assets/{file:[\w\W\/]+}`, func(wr http.ResponseWriter, r *http.Request) {
		http.ServeFile(wr, r, r.URL.Path[1:])
	})

	router.HandleFunc("/game/socket/", game.SocketHandler)

	http.ListenAndServe(":8080", router)
}

func main() {
	SetupRoutes()
}
