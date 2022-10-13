package types

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
	Soldiers   []*Soldier
	Collisions []Collision
	Radius     Map2D
}

func (w *World) Update() {
	for _, soldier := range w.Soldiers {
		soldier.Update()
	}

}

func (w *World) NewSoldier(id int, faction string, position Map2D) *Soldier {
	s := &Soldier{
		Faction: faction,
		Health:  MaxHealth,

		Position: position,
		Velocity: Map2D{},
		Radius:   Map2D{X: 0.5, Y: 1},

		PointOfShooting: 0.6,
		ReloadingSpeed:  ReloadingSpeed,
		LastShot:        0,

		WeaponDamage: MaxDamage,
		Direction:    true,
		Id:           id,

		World: w,
	}

	w.Soldiers = append(w.Soldiers, s)
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
