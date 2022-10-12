package controller

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/backend/game"
	"golang.org/x/net/websocket"
)

var (
	g = game.NewGame("world.json", game.NewSoldier(10, 10, 5, 0, 0, false), game.NewSoldier(10, 10, 5, 0, 0, false))
)

func SendInfo(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(c *websocket.Conn) {

		for {
			// use the info for getting the soldiers and all that stuff
			websocket.JSON.Send(c, g)
			// just use the data that you have for seeing it in the game
		}

	}).ServeHTTP(w, r)
}
