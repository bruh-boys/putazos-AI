package types

import (
	"golang.org/x/net/websocket"
)

const (
	FramesPerSecond = 30.0
	Gravity         = 1
)

type Map2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Collision struct {
	Position Map2D
	Radius   Map2D
}

type World struct {
	Soldiers    map[*websocket.Conn]*Soldier
	Collisions  []Collision
	Radius      Map2D
	Initialized bool
}

func (w *World) Update() {
	for _, soldier := range w.Soldiers {
		soldier.Update()
	}

}

var Factions = map[bool]string{
	false: "blue",
	true:  "red",
}

func (w *World) NewSoldier(ws *websocket.Conn) *Soldier {
	s := &Soldier{
		Faction: Factions[len(w.Soldiers)%2 == 0],
		Health:  MaxHealth,

		Position: Map2D{X: 0, Y: 0},
		Velocity: Map2D{X: 0, Y: 0},
		Radius:   Map2D{X: 0.5, Y: 1},

		PointOfShooting: 0.6,
		ReloadingSpeed:  ReloadingSpeed,
		LastShot:        0,

		WeaponDamage: MaxDamage,
		Direction:    true,

		World: w,
	}

	w.Soldiers[ws] = s

	return s
}

func (w *World) GenerateEntity(pos Map2D, entity interface{}) {

}

func (w World) OnCollision(pos Map2D, radius Map2D) (Collision, bool) {

	for _, collision := range w.Collisions {
		if pos.X+radius.X > collision.Position.X &&
			pos.X < collision.Position.X+collision.Radius.X &&
			pos.Y+radius.Y > collision.Position.Y &&
			pos.Y < collision.Position.Y+collision.Radius.Y {

			return collision, true
		}

	}

	return Collision{}, false
}
