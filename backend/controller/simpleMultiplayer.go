package controller

import (
	"net/http"

	"github.com/bruh-boys/putazos-ai/backend/game"
	"golang.org/x/net/websocket"
)

type Player struct {
	Soldier *game.Soldier
	Id      int
}

var (
	Connections   = map[*websocket.Conn]Player{}
	games         = map[int]map[int]*game.Soldier{}
	g             = game.NewGame("world.json")
	width, height = 27.0, 31.0
	id            = 0
	last          = 0
)

// really basic multiplayer implemented with the ass
func Multiplayer(w http.ResponseWriter, r *http.Request) {
	s := websocket.Server{
		Handler: websocket.Handler(func(c *websocket.Conn) {
			play := Player{
				(game.NewSoldier("Blue",
					width, height,
					15,
					0,
					0,
					true,
				)),
				id % 2}
			Connections[c] = play

			if id/2 != last || len(games) == 0 {
				games[id/2] = map[int]*game.Soldier{}

				id = last
				last = id
			}

			games[id/2][play.Id] = Connections[c].Soldier

			g.Spawn(Connections[c].Soldier)
			identification := id / 2
			id++
			go func() {
				for {
					action := ""
					if err := websocket.Message.Receive(c, &action); err != nil {
						return
					}
					g.Action(Connections[c].Id, action, games[identification])

				}
			}()

			for {
				g.DoSomethingPerFrame(play.Soldier)
				if err := websocket.Message.Send(c, games[identification]); err != nil {
					delete(Connections, c)
					delete(games[identification], play.Id)
					if len(games[identification]) == 0 {
						delete(games, identification)
					}
				}
			}

		})}

	s.ServeHTTP(w, r)
}
