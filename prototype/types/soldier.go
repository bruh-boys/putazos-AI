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

	s.World.GenerateEntity(Map2D{
		Y: s.Position.Y + (s.Radius.Y / 2),
		X: s.Position.X + s.Radius.X,
	}, "bullet")

}

func (s *Soldier) Move() {
	s.Velocity.Y += Gravity

	s.Position.X += (s.Velocity.X / FramesPerSecond)
	s.Position.Y += (s.Velocity.Y / FramesPerSecond)

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

var actions = map[string]func(s *Soldier){
	"left": func(s *Soldier) {
		s.Velocity.X = -1
	},
	"right": func(s *Soldier) {
		s.Velocity.X = 1
	},
	"up": func(s *Soldier) {
		s.Velocity.Y = 1.5
	},
	"down": func(s *Soldier) {

	},
}

func (s *Soldier) Update() bool {
	for action, value := range s.Actions {
		if value {
			actions[action](s)
		}
	}

	s.Move()

	if s.Health <= 0 && s.Velocity.Y == 0 {

		return true
	}

	if s.Health > 0 && s.Actions["shoot"] {
		s.Shoot()

	}

	return false
}
