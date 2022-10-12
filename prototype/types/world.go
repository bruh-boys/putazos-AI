package prototype

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
	Collisions []Collision
	Radius     Map2D
}

func (w World) GenerateEntity(pos Map2D, entity string) {

}

func (w World) OnCollision(pos Map2D, radius Map2D) (Collision, bool) {

	for _, collision := range w.Collisions {
		if pos.X+radius.X >= collision.Position.X &&
			pos.X <= collision.Position.X+collision.Radius.X &&
			pos.Y+radius.Y >= collision.Position.Y &&
			pos.Y <= collision.Position.Y+collision.Radius.Y {

			return collision, true
		}

	}

	return Collision{}, false
}
