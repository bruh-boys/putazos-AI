package game

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/models"
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

type Request struct {
	Action string `json:"action"`
	Active bool   `json:"active"`
	Data   string
}

func ListenMessages(ws *websocket.Conn) {
	for {
		var request Request

		if err := websocket.JSON.Receive(ws, &request); err != nil {

			continue
		}

		models.Game[ws].Soldiers[ws].Action(
			request.Action, request.Active,
		)

	}

}

func SendMessages(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		models.Game[ws].Update()

		var (
			soldiers    []types.Soldier    = []types.Soldier{}
			projectiles []types.Projectile = []types.Projectile{}
		)

		for _, soldier := range models.Game[ws].Soldiers {
			soldiers = append(soldiers, types.Soldier{
				Faction:     soldier.Faction,
				Health:      soldier.Health,
				Position:    soldier.Position,
				IsCrouching: soldier.IsCrouching,
			})
		}

		for _, projectile := range models.Game[ws].Projectiles {
			projectiles = append(projectiles, types.Projectile{
				Position: projectile.Position,
			})
		}

		var response = &struct {
			Soldiers    []types.Soldier    `json:"soldiers"`
			Projectiles []types.Projectile `json:"projectiles"`
		}{
			Soldiers:    soldiers,
			Projectiles: projectiles,
		}

		v, _ := json.Marshal(response)

		websocket.Message.Send(ws, string(v))

	}

}

func SocketHandler(wr http.ResponseWriter, r *http.Request) {
	socket := websocket.Server{
		Handler: func(ws *websocket.Conn) {
			models.NewGame(ws)

			go ListenMessages(ws)
			SendMessages(ws)
		},
	}

	socket.ServeHTTP(wr, r)
}
