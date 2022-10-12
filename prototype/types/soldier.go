package prototype

import (
	"time"
)

const (
	RateFire  = 500
	MaxHealth = 100
	MaxDamage = 15
	MaxAmmo   = 30
)

type Soldier struct {
	Faction string `json:"faction"`
	Health  uint8  `json:"health"`

	Position Map2D `json:"position"`
	Velocity Map2D
	Radius   Map2D `json:"radius"`

	PointOfShooting float64
	ReloadingSpeed  float64
	LastShot        int64

	WeaponDamage uint8
	Direction    bool

	World *World
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
	// I think with math you only need calculate the collision once.
	// But I don't know how to do it.

	if _, on := s.World.OnCollision(Map2D{
		X: s.Position.X + (s.Velocity.X / FramesPerSecond),
		Y: s.Position.Y,
	}, s.Radius); on {
		s.Velocity.X = 0
	}

	if _, on := s.World.OnCollision(Map2D{
		Y: s.Position.Y + (s.Velocity.Y / FramesPerSecond),
		X: s.Position.X,
	}, s.Radius); on {
		s.Velocity.Y = 0
	}

}

func (s *Soldier) Update() bool {
	s.Move()

	if s.Health <= 0 && s.Velocity.Y == 0 {

		return true
	}

	if s.Health > 0 {
		s.Shoot()

	}

	return false
}
