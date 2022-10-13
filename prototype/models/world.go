package models

import (
	"time"

	"github.com/bruh-boys/putazos-ai/prototype/types"
	"golang.org/x/net/websocket"
)

var World = &types.World{
	Soldiers:    map[*websocket.Conn]*types.Soldier{},
	Collisions:  []types.Collision{},
	Initialized: false,
}

func RunWorld() *types.World {
	World.Initialized = true

	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		World.Update()
	}
}
