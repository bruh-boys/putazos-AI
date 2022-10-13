package game

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/models"
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

type Response struct {
	Action string `json:"action"`
	Active bool   `json:"active"`
}

func ListenMessages(ws *websocket.Conn) {
	for {
		var response Response

		if err := websocket.Message.Receive(ws, &response); err != nil {

			continue
		}

		models.World.Soldiers[ws].Action(
			response.Action, response.Active,
		)

	}

}

func SendMessages(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)

		var soldiers []types.Soldier

		for _, soldier := range models.World.Soldiers {
			soldiers = append(soldiers, *soldier)
		}

		websocket.JSON.Send(ws, fmt.Sprintf("%#v", soldiers))

	}

}

func SocketHandler(wr http.ResponseWriter, r *http.Request) {
	socket := websocket.Server{
		Handler: func(ws *websocket.Conn) {
			models.World.NewSoldier(ws)

			if models.World.Initialized == false {
				go models.RunWorld()

			}

			go ListenMessages(ws)

			SendMessages(ws)
		},
	}

	socket.ServeHTTP(wr, r)
}
