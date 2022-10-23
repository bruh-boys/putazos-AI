package game

import (
	"net/http"
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/models"
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

type GameDataProjectile struct {
	Position types.Map2D `json:"position"`
	Start    types.Map2D `json:"start"`
}

type GameDataSoldier struct {
	Direction bool        `json:"direction"`
	Faction   string      `json:"faction"`
	Health    uint8       `json:"health"`
	Position  types.Map2D `json:"position"`
	Actions   []string    `json:"actions"`
	Id        string      `json:"id"`
}

type GameData struct {
	Projectiles []GameDataProjectile `json:"projectiles"`
	Soldiers    []GameDataSoldier    `json:"soldiers"`
}

type GameResponse struct {
	Type string   `json:"type"`
	Data GameData `json:"data"`
}

type JoinResponse struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Request struct {
	Type string `json:"type"`
	Data struct {
		State bool   `json:"state"`
		Key   string `json:"key"`
	} `json:"data"`
}

func ListenMessages(ws *websocket.Conn) {
	for {
		var request Request

		if err := websocket.JSON.Receive(ws, &request); err != nil {

			continue
		}

		switch request.Type {
		case "action":
			models.Game[ws].Soldiers[ws].Action(
				request.Data.Key, request.Data.State,
			)
		}

	}

}

func SendMessages(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second / types.FramesPerSecond)
		models.Game[ws].Update()

		var response GameResponse = GameResponse{
			Type: "update",
			Data: GameData{},
		}

		for _, projectile := range models.Game[ws].Projectiles {
			response.Data.Projectiles = append(response.Data.Projectiles, GameDataProjectile{
				Position: projectile.Position,
				Start:    projectile.Start,
			})
		}

		for _, soldier := range models.Game[ws].Soldiers {
			var actions []string

			for action, state := range soldier.Actions {
				if state {
					actions = append(actions, action)
				}
			}

			response.Data.Soldiers = append(response.Data.Soldiers, GameDataSoldier{
				Faction:   soldier.Faction,
				Health:    soldier.Health,
				Position:  soldier.Position,
				Actions:   actions,
				Id:        soldier.Id,
				Direction: soldier.Direction,
			})

		}

		websocket.JSON.Send(ws, response)
	}

}

func SocketHandler(wr http.ResponseWriter, r *http.Request) {
	socket := websocket.Server{
		Handler: func(ws *websocket.Conn) {
			data := models.NewGame(ws)

			go ListenMessages(ws)

			websocket.JSON.Send(ws, JoinResponse{
				Type: "join",
				Data: string(data),
			})

			SendMessages(ws)
		},
	}

	socket.ServeHTTP(wr, r)
}
