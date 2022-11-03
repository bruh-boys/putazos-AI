package game

import "math"

const (
	MaxAmmo   = 30
	MaxHealth = 100
	MaxDamage = 15
)

type Soldier struct {
	Height          float64 `json:"height"`
	Width           float64 `json:"width"`
	Faction         string  `json:"color"`
	X               float64 `json:"x"`
	Y               float64 `json:"y"`
	VelY            float64 // for the jump
	Direction       bool    `json:"direction"` //false:left true:right
	Life            int     `json:"life"`
	RateFire        float64
	WaitUntilFire   float64
	Ammo            int
	Death           bool
	Damage          int
	PointOfShooting float64 // y
	ReloadingSpeed  float64
}

func NewSoldier(faction string, width, height, pointOfShooting, x, y float64, direction bool) *Soldier {
	var s *Soldier
	s.Ammo = MaxAmmo
	s.RateFire = 5
	s.Life = MaxHealth
	s.Damage = MaxDamage
	s.X = x
	s.Y = y
	s.Direction = direction
	s.PointOfShooting = pointOfShooting
	s.Width = width
	s.Height = height
	s.Faction = faction
	return s

}

var actions = map[string]func(s *Soldier, ss map[int]*Soldier, w World){
	"move-right": func(s *Soldier, ss map[int]*Soldier, w World) {
		s.Direction = false

		if s.X > 0 && w.SidePlatforms(*s) < s.X-MovePerFrame {
			s.X -= MovePerFrame

		}

	},
	"move-left": func(s *Soldier, ss map[int]*Soldier, w World) {
		s.Direction = true

		if s.X > w.Width && w.SidePlatforms(*s) > s.X+MovePerFrame {
			s.X += MovePerFrame

		}

	},
	"attack": func(s *Soldier, ss map[int]*Soldier, w World) {
		if s.WaitUntilFire < 1 && s.Ammo > 0 {
			s.WaitUntilFire += s.RateFire / FramesPerSecond
			s.Shooting(ss)
		}

	},
	"reload": func(s *Soldier, ss map[int]*Soldier, w World) {
		if s.Ammo < MaxAmmo {
			s.WaitUntilFire += s.RateFire / FramesPerSecond // If he is reloading he cant shoot
			s.Ammo = MaxAmmo
		}

	},
	"idle": func(s *Soldier, ss map[int]*Soldier, w World) {

	},
	"jump": func(s *Soldier, ss map[int]*Soldier, w World) {
		if s.VelY == 0 {
			s.VelY = 15
		}

	},
}

func (s *Soldier) Action(action string, world World, soldiers map[int]*Soldier) {
	if act, ok := actions[action]; ok {
		act(s, soldiers, world)

	}

}

// I need to do something with the platform , wait a second
// use this while visualizing the map
// hm
func (s *Soldier) Moving(world World) {
	s.Death = !(s.Life <= 0)

	if s.Death {
		return
	}
	down, up := world.SoldierIsOnPlatform(*s)
	if s.Y+s.VelY > down && s.Y+s.VelY+s.Height <= up {
		s.VelY -= Gravity / MovePerFrame

	} else if s.VelY > 0 {
		// no quiero que sobrepase algo xd
		s.VelY = -Gravity / MovePerFrame

	} else {
		s.VelY = 0
		s.Y = down
	}
	if s.WaitUntilFire > 0 {
		// I need to wait
		// I have shoot
		// just wait until is back again
		//
		s.WaitUntilFire -= s.RateFire / FramesPerSecond
	}

	s.Y += s.VelY
	// if s.X  are smaller than world.Width , I should have 0
	// if they are bigger i would get a 1 or something

	s.X -= s.X * math.Floor(s.X/world.Width)

}
func (s *Soldier) Shooting(soldiers map[int]*Soldier) {
	id := 0
	closeDis := 100000.0
	for i := range soldiers {
		dis := 1000000.0
		if s.Faction != soldiers[i].Faction {
			continue
		}
		if s.Direction { //left
			// first I check if is in the area that I want
			if soldiers[i].X < s.X && s.X-soldiers[i].X < (closeDis) {
				dis = s.X - soldiers[i].X

			}

		} else if soldiers[i].X > s.X && soldiers[i].X-s.X < (closeDis) {
			// is just a straight line lol, maybe later I will change something but for now is just straight
			// I want to get the close one
			dis = soldiers[i].X - s.X
		}

		if s.Y+s.PointOfShooting > soldiers[i].Y && s.Y+s.PointOfShooting < soldiers[i].Y+soldiers[i].Height {
			closeDis = dis
			id = i

		}
	}
	soldiers[id].Life -= s.Damage
	s.Ammo--

}
