package models

import (
	"encoding/json"
	"os"

	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

var Game = map[*websocket.Conn]*types.World{}

type Map2 struct {
	Defs []struct {
		Positions []types.Map2D `json:"positions"`
		Id        string        `json:"id"`
	} `json:"defs"`
}

func NewGame(ws *websocket.Conn) {
	var m Map2

	data, _ := os.ReadFile("assets/map.json")

	json.Unmarshal(data, &m)

	Game[ws] = &types.World{
		Soldiers: map[*websocket.Conn]*types.Soldier{},
	}

	for _, def := range m.Defs {
		if def.Id == "spawn-point" {
			Game[ws].SpawnPoints = def.Positions
		} else if def.Id == "collision" {
			for _, position := range def.Positions {
				Game[ws].Collisions = append(Game[ws].Collisions, types.Collision{
					Position: position,
					Radius:   types.Map2D{X: 16, Y: 16},
				})
			}
		}
	}

	Game[ws].NewSoldier(ws)
}

/*
func RunWorld() *types.World {
	World.Initialized = true

	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		World.Update()
	}
}
*/
