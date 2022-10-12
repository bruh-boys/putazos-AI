package game

import (
	"encoding/json"
	"os"
	"time"
)

type Game struct {
	World       World      `json:"world"`
	Soldiers    []Soldier  `json:"soldiers"`
	SpawnPoints []Platform `json:"spawn-points"`
	GameMap     []float64  `json:"game-map"`
}

func NewGame(name string, blue, red Soldier) (g Game) {
	file, _ := os.Open(name)
	json.NewDecoder(file).Decode(&g)
	g.Soldiers = []Soldier{blue, red}
	return
}

// Maybe I should do this in at the same time idk
func (g Game) Action(id int, action string) {
	g.Soldiers[id].Action(action, g.World, g.Soldiers)

}
func (g Game) DoSomethingPerFrame() {
	for _, s := range g.Soldiers {
		s.Moving(g.World)
	}
	time.Sleep(time.Second / FramesPerSecond)
}
