package controller

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/backend/game"
	"golang.org/x/net/websocket"
)

type Player struct {
	Soldier game.Soldier
	Id      int
}

var (
	Connections   = map[*websocket.Conn]Player{}
	games         = [][]game.Soldier{}
	g             = game.NewGame("world.json")
	width, height = 27.0, 31.0
	id            = 0
	last          = 0
)

// really basic multiplayer implemented with the ass
func Multiplayer(w http.ResponseWriter, r *http.Request) {
	s := websocket.Server{
		Handler: websocket.Handler(func(c *websocket.Conn) {
			Connections[c] = Player{
				game.NewSoldier(map[bool]string{
					true:  "blue",
					false: "red",
				}[id%2 == 0],
					width, height,
					15,
					0,
					0,
					true,
				), id % 2}

			if id/2 != last || len(games) == 0 {
				games = append(games, []game.Soldier{
					Connections[c].Soldier,
				})

				id = last
			}

			games[id/2] = append(games[id/2], Connections[c].Soldier)
	
			g.Spawn(Connections[c].Soldier)
			identification := id / 2
			id++
			go func() {
				for {
					action := ""
					websocket.Message.Receive(c, &action)
					g.Action(Connections[c].Id, action, games[identification])

				}
			}()

			for {
				g.DoSomethingPerFrame(games[identification])
				websocket.Message.Send(c, games[identification])
			}

		})}

	s.ServeHTTP(w, r)
}
