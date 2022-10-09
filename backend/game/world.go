package game

var (
	MovePerFrame    = 10.0
	FramesPerSecond = 30
	Gravity         = 1.0
)

type Platform struct {
	X      float64
	Y      float64
	Base   float64
	Height float64
}

type World struct {
	Width, Height float64
	Platforms     []Platform
}

func (w World) SoldierIsOnPlatform(soldier Soldier) (float64, float64) {

	down := 0.0
	downDis := 0.0
	lowDown := 100000.0
	lowDisDown := 100000.0

	up := 0.0
	upDis := 0.0
	lowUp := 100000.0
	lowDisUp := 100000.0

	for i := 0; i < len(w.Platforms); i++ {

		if soldier.Y <= (w.Platforms[i].Y) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Base) && soldier.X >= w.Platforms[i].X) {
			upDis = soldier.Y - (w.Platforms[i].Y + w.Platforms[i].Height)
			up = w.Platforms[i].Y
		}
		if lowDisUp > upDis && upDis >= 0 {
			lowDisUp = upDis
			lowUp = up
		}

		// is up or down?
		if soldier.Y >= (w.Platforms[i].Y+w.Platforms[i].Height) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Base) && soldier.X >= w.Platforms[i].X) {
			downDis = soldier.Y - (w.Platforms[i].Y + w.Platforms[i].Height)
			down = w.Platforms[i].Y + w.Platforms[i].Height
		}
		if lowDisDown > downDis && downDis >= 0 {
			lowDown = down
			lowDisDown = downDis
		}

	}
	return lowDown, lowUp
}
