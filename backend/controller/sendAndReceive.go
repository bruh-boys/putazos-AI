package controller

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/websocket"
)

func SendInfo(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(c *websocket.Conn) {

		for {
			// use the info for getting the soldiers and all that stuff
			websocket.JSON.Send(c, []float64{})
			// just use the data that you have for seeing it in the game
		}

	}).ServeHTTP(w, r)
}
func SeeGame(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(g)
}
