package types

type Projectile struct {
	Position Map2D
	Radius   Map2D
	Velocity Map2D
	World    *World
}

func (p *Projectile) Update() {
	p.Velocity.Y += Gravity / FramesPerSecond

	p.Position.X += p.Velocity.X / FramesPerSecond
	p.Position.Y += p.Velocity.Y / FramesPerSecond

	if coll, on := p.World.OnCollision(p.Position, p.Radius); on {
		if p.Velocity.X < 0 {
			p.Position.X = coll.Position.X + coll.Radius.X + 0.01

		}

		if p.Velocity.X > 0 {
			p.Position.X = coll.Position.X - coll.Radius.X - 0.01

		}

	}

	if coll, on := p.World.OnCollision(p.Position, p.Radius); on {
		if p.Velocity.Y < 0 {
			p.Velocity.Y = 0
			p.Position.Y = coll.Position.Y + coll.Radius.Y + 0.01
		}

		if p.Velocity.Y > 0 {
			p.Velocity.Y = 0
			p.Position.Y = coll.Position.Y - coll.Radius.Y - 0.01

		}

	}

}
