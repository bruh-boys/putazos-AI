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

func (w World) OnCollision(pos Map2D, radius float64) (Map Map2D) {
	for _, collision := range w.Collisions {

		// I think it's working, but i'm not sure.
		if pos.X+radius >= collision.Position.X &&
			pos.X <= collision.Position.X+collision.Radius.X {
			return Map2D{
				X: collision.Position.X,
				Y: -0.01,
			}

		}

		if pos.Y+radius >= collision.Position.Y &&
			pos.Y <= collision.Position.Y+collision.Radius.Y {
			return Map2D{
				Y: collision.Position.Y,
				X: -0.01,
			}

		}

	}

	return Map2D{}
}
