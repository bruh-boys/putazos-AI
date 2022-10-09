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
	lowy := 10000.0
	lowDis := 100000.0
	y := 0.0
	dis := 0.0

	for i := 0; i < len(w.Platforms); i++ {
		// is up or down?
		if soldier.Y >= (w.Platforms[i].Y+w.Platforms[i].Height) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Base) && soldier.X >= w.Platforms[i].X) {
			dis = soldier.Y - (w.Platforms[i].Y + w.Platforms[i].Height)
			y = w.Platforms[i].Y + w.Platforms[i].Height
		}
		if lowDis > dis && dis >= 0 {
			lowy = y
			lowDis = dis
		}

	}
	return lowy
}
