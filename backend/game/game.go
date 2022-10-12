package game

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	World       World      `json:"world"`
	SpawnPoints []Platform `json:"spawn-points"`
	GameMapO    []float64  `json:"game-map"`
}

func NewGame(name string) (g Game) {
	file, _ := os.Open(name)
	json.NewDecoder(file).Decode(&g)

	return
}

// Maybe I should do this in at the same time idk
func (g Game) Action(id int, action string, soldiers []Soldier) {
	soldiers[id].Action(action, g.World, soldiers)

}
func (g Game) Spawn(soldier Soldier) {
	soldier.Life = MaxHealth
	rand.Shuffle(len(g.SpawnPoints), func(i, j int) { g.SpawnPoints[i], g.SpawnPoints[j] = g.SpawnPoints[j], g.SpawnPoints[i] })
	soldier.X = g.SpawnPoints[1].X
	soldier.Y = g.SpawnPoints[1].Y

}
func (g Game) DoSomethingPerFrame(soldiers []Soldier) {

	for _, s := range soldiers {
		s.Moving(g.World)
		if s.Death {
			g.Spawn(s)
		}
	}
	time.Sleep(time.Second / FramesPerSecond)
}
