package game

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/models"
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

var (
	Connections = map[*websocket.Conn]*types.Soldier{}
)

var Factions = map[bool]string{
	false: "blue",
	true:  "red",
}

func ListenMessages(ws *websocket.Conn) {
	for {
		var action string = ""

		if err := websocket.Message.Receive(ws, &action); err != nil {

			continue
		}

		act := strings.Split(strings.ReplaceAll(action, " ", ""), "@")

		if val, err := strconv.ParseBool(act[1]); err == nil {
			Connections[ws].Action(act[0], val)

		}

	}

}

func SendMessages(ws *websocket.Conn) {
	id := len(Connections) + 1

	Connections[ws] = models.NewSoldier(
		id,
		Factions[id%2 == 0],
		types.Map2D{
			X: 0,
			Y: 0,
		},
	)

	for {
		Connections[ws].World.Update()

		websocket.Message.Send(ws, struct {
			Soldiers []*types.Soldier
		}{
			Soldiers: Connections[ws].World.Soldiers,
		})

	}

}

func SocketHandler(wr http.ResponseWriter, r *http.Request) {
	socket := websocket.Server{
		Handler: func(ws *websocket.Conn) {
			go ListenMessages(ws)

			for {
				time.Sleep(time.Second * 2)

				SendMessages(ws)
			}
		},
	}

	socket.ServeHTTP(wr, r)
}
