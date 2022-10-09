package game

var (
	MovePerFrame    = 10.0
	FramesPerSecond = 30
	Gravity         = 1.0
)

type Soldier struct {
	Height    int
	Widht     int
	Faction   bool
	X         float64
	Y         float64
	VelY      float64 // for the jump
	Direction bool    //false:left true:right
	Life      int
}

// just a simple rectangle
// the begining is on the left down corner

func (s *Soldier) Action(action string, world World) {
	switch action {
	case "move-left":
		s.X -= MovePerFrame
		s.Direction = false
	case "move-right":
		s.X += MovePerFrame
		s.Direction = true
	case "jump":
		// in the world this should
		s.VelY += 15

	case "shoot":
		// use this for the world

	}
}

// I need to do something with the platform , wait a second
func (s *Soldier) Moving(world World) {
	s.Y += s.VelY

	dpy := world.SoldierIsOnPlatform(*s)
	if s.Y > dpy {
		s.VelY -= Gravity

		s.X += map[bool]float64{true: MovePerFrame / 2, false: -MovePerFrame / 2}[s.Direction]
	} else {
		s.VelY = 0
		s.Y = dpy
	}

}
