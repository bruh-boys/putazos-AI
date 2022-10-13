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
	Games       = map[int]*types.World{}
)

var Factions = map[bool]string{
	false: "blue",
	true:  "red",
}

func NewGame(c chan int) {
	id := len(Games) + 1

	Games[id] = &types.World{}

	c <- id
	<-c

	for {
		time.Sleep(time.Second / types.FramesPerSecond)

		if len(Games[id].Soldiers) == 0 {

			delete(Games, id)

			break
		}

		for _, soldier := range Connections {
			soldier.Move()

		}

	}

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
	var world *types.World = nil

	for _, game := range Games {
		if len(game.Soldiers) < 6 {
			world = game

		}

	}

	c := make(chan int)

	if world == nil {
		go NewGame(c)

	}

	id := <-c

	world.Soldiers = append(world.Soldiers, models.NewSoldier(
		id,
		Factions[id%2 == 0],
		types.Map2D{
			X: 0,
			Y: 0,
		},
	))

	c <- -1

	for {
		time.Sleep(time.Second / types.FramesPerSecond)

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
