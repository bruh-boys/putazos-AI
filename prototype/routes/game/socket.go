package game

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/models"
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

type Request struct {
	Action string `json:"action"`
	Active bool   `json:"active"`
}

func ListenMessages(ws *websocket.Conn) {
	for {
		var request Request

		if err := websocket.Message.Receive(ws, &request); err != nil {

			continue
		}

		models.Game[ws].Action(
			request.Action, request.Active,
		)

	}

}

func SendMessages(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		var soldiers []types.Soldier

		for _, soldier := range models.Game[ws].World.Soldiers {
			soldiers = append(soldiers, *soldier)
		}

		websocket.JSON.Send(ws, fmt.Sprintf("%#v", soldiers))

	}

}

func SocketHandler(wr http.ResponseWriter, r *http.Request) {
	socket := websocket.Server{
		Handler: func(ws *websocket.Conn) {
			models.Game[ws].World.NewSoldier(ws)

			go ListenMessages(ws)
			SendMessages(ws)
		},
	}

	socket.ServeHTTP(wr, r)
}
