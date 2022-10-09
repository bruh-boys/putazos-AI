package game

var (
	MovePerFrame    = 10.0
	FramesPerSecond = 30
	Gravity         = 1.0
)

type Platform struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Base   float64 `json:"base"`
	Height float64 `json:"height"`
}

type World struct {
	Width     float64    `json:"width"`
	Height    float64    `json:"height"`
	Platforms []Platform `json:"platforms"`
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
		// with this i get the platform that its up
		if soldier.Y <= (w.Platforms[i].Y) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Base) && soldier.X >= w.Platforms[i].X) {
			upDis = (w.Platforms[i].Y) - soldier.Y
			up = w.Platforms[i].Y
		}
		if lowDisUp > upDis && upDis >= 0 {
			lowDisUp = upDis
			lowUp = up
		}
		// the platform that is down
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
