package prototype

import "time"

const (
	RateFire  = 500
	MaxHealth = 100
	MaxDamage = 15
	MaxAmmo   = 30
)

type Soldier struct {
	Faction string  `json:"faction"`
	Radius  float64 `json:"radius"`
	Health  uint8   `json:"health"`

	Position Map2D `json:"position"`
	Velocity Map2D

	PointOfShooting float64
	ReloadingSpeed  float64
	LastShot        int64

	WeaponDamage uint8
	Direction    bool

	World *World
}

func (s *Soldier) Shoot() {
	if (s.LastShot + RateFire) > time.Now().UnixNano() {

		return
	}

	s.LastShot = time.Now().UnixMilli()

	s.World.GenerateEntity(Map2D{
		Y: s.Position.Y + (s.Radius / 2),
		X: s.Position.X + s.Radius,
	}, "bullet")

}

func (s *Soldier) Move() {
	// Check if the next position is valid.
	world := s.World.OnCollision(Map2D{
		X: s.Position.X + (s.Velocity.X / FramesPerSecond),
		Y: s.Position.Y + (s.Velocity.Y / FramesPerSecond),
	}, s.Radius)

	// If the next position is not valid, move the soldier to the valid position.
	// If the next position is valid, move the soldier to the next position.
	if world.X != -0.01 {
		if s.Direction {
			s.Position.X += s.Velocity.X / FramesPerSecond

		} else {
			s.Position.X -= s.Velocity.X / FramesPerSecond

		}

	} else {
		s.Velocity.X = world.X

	}

	// If the next position is not valid, move the soldier to the valid position.
	// If the next position is valid, move the soldier to the next position.
	if world.Y != -0.01 {
		// Doing this create a jump effect.
		if s.Velocity.Y > 0 {
			s.Position.Y += s.Velocity.Y / FramesPerSecond

		} else {
			s.Position.Y -= s.Velocity.Y / FramesPerSecond

		}

		s.Velocity.Y -= Gravity / FramesPerSecond
	} else {
		// Gravity damage.
		if s.Velocity.Y < 0 {
			s.Health += uint8(s.Velocity.Y / -4)

		}

		s.Position.Y = world.Y
		s.Velocity.Y = 0

	}

}

func (s *Soldier) Update() bool {
	s.Move()

	if s.Health <= 0 && s.Velocity.Y == 0 {

		return true
	}

	s.Shoot()

	return false
}
