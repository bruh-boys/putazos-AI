package models

import (
	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

var Game = map[*websocket.Conn]*types.Soldier{}

/*
func RunWorld() *types.World {
	World.Initialized = true

	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		World.Update()
	}
}
*/
