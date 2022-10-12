package game

import (
	"encoding/json"
	"fmt"
	"os"
)

type Game struct {
	World       World      `json:"world"`
	Soldiers    []Soldier  `json:"soldiers"`
	SpawnPoints []Platform `json:"spawn-points"`
	GameMap     []float64  `json:"game-map"`
}

func NewGame(name string) (g Game) {
	file, _ := os.Open(name)
	json.NewDecoder(file).Decode(&g)
	return
}
func DoSomething() {
	fmt.Println("idk ")
}
