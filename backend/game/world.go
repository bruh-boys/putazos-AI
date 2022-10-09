package game

type Platform struct {
	X      float64
	Y      float64
	Base   float64
	Height float64
}
type World struct {
	Platforms []Platform
	Soldiers  []Soldier
}

func (w World) SoldierIsOnPlatform(soldier Soldier) float64 {
	y := 0.0
	for i := 0; i < len(w.Platforms); i++ {

		// is up or down?
		if soldier.Y >= (w.Platforms[i].Y+w.Platforms[i].Height) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Base) && soldier.X >= w.Platforms[i].X) {
			y = w.Platforms[i].Y + w.Platforms[i].Height
		}

	}
	return y
}
