package types

import (
	"time"
)

const (
	ReloadingSpeed = 1500
	RateFire       = 500

	MaxHealth uint8 = 100
	MaxDamage uint8 = 15
	MaxAmmo   uint8 = 30
)

type Soldier struct {
	Faction string `json:"faction"`
	Health  uint8  `json:"health"`

	Position Map2D `json:"position"`
	Velocity Map2D
	Radius   Map2D `json:"radius"`

	Actions map[string]bool

	PointOfShooting float64
	ReloadingSpeed  float64
	LastShot        int64

	IsCrouching bool

	WeaponDamage uint8
	Direction    bool

	World *World
	Id    int
}

func (s *Soldier) Action(action string, value bool) {
	s.Actions[action] = value

}

func (s *Soldier) Shoot() {
	if s.LastShot > time.Now().UnixNano() {

		return
	}

	s.LastShot = time.Now().UnixMilli() + RateFire

	if s.Direction {
		s.World.GenerateEntity(Map2D{
			X: s.Position.X + s.PointOfShooting,
			Y: s.Position.Y,
		})
	} else {
		s.World.GenerateEntity(Map2D{
			X: s.Position.X - s.PointOfShooting,
			Y: s.Position.Y,
		})
	}
}

func (s *Soldier) Move() {
	s.Velocity.Y += Gravity / FramesPerSecond

	s.Position.X += s.Velocity.X / FramesPerSecond
	s.Position.Y += s.Velocity.Y / FramesPerSecond

	if coll, on := s.World.OnCollision(s.Position, s.Radius); on {
		if s.Velocity.X < 0 {
			s.Position.X = coll.Position.X + coll.Radius.X + 0.01

		}

		if s.Velocity.X > 0 {
			s.Position.X = coll.Position.X - coll.Radius.X - 0.01

		}

	}

	if coll, on := s.World.OnCollision(s.Position, s.Radius); on {
		if s.Velocity.Y < 0 {
			s.Velocity.Y = 0
			s.Position.Y = coll.Position.Y + coll.Radius.Y + 0.01
		}

		if s.Velocity.Y > 0 {
			s.Velocity.Y = 0
			s.Position.Y = coll.Position.Y - coll.Radius.Y - 0.01

		}

	}

}

var actions = map[string]func(s *Soldier, b bool){
	"left": func(s *Soldier, b bool) {
		if b && s.IsCrouching {
			s.Velocity.X = -1
		} else if b {
			s.Velocity.X = -2
		} else {
			s.Velocity.X = 0
		}
	},
	"right": func(s *Soldier, b bool) {
		if b && s.IsCrouching {
			s.Velocity.X = 1
		} else if b {
			s.Velocity.X = 2
		} else {
			s.Velocity.X = 0
		}
	},
	"up": func(s *Soldier, b bool) {
		if b && s.Velocity.Y == 0 {
			s.Velocity.Y = -1
		}
	},
	"down": func(s *Soldier, b bool) {
		if b && s.Velocity.Y == 0 {
			s.IsCrouching = true
		} else {
			s.IsCrouching = false
		}
	},
}

func (s *Soldier) Update() bool {
	if s.Health <= 0 && s.Velocity.Y == 0 {

		return true
	}

	for action, value := range s.Actions {
		actions[action](s, value)
	}

	s.Move()

	if s.Health > 0 && s.Actions["shoot"] {
		s.Shoot()

	}

	return false
}
