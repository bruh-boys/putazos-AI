package game

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
	down, _ := world.SoldierIsOnPlatform(*s)
	switch action {
	case "move-left":
		s.X -= MovePerFrame
		s.Direction = false
	case "move-right":

		s.X += MovePerFrame
		s.Direction = true
	case "jump":
		// in the world this should

		if s.Y == down {
			s.VelY = 15
		}

	case "shoot":
		// use this for the world

	}
}

// I need to do something with the platform , wait a second
func (s *Soldier) Moving(world World) {
	down, up := world.SoldierIsOnPlatform(*s)
	if s.Y+s.VelY > down && s.Y+s.VelY <= up {
		s.VelY -= Gravity

	} else if s.Y+s.VelY >= up && s.VelY > 0 {
		s.VelY = -Gravity

	} else {
		s.VelY = 0
		s.Y = down
	}
	s.Y += s.VelY

}
