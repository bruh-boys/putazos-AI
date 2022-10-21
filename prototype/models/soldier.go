package models

import "github.com/bruh-boys/putazos-ai/prototype/types"

func NewSoldier(faction string, position types.Map2D) *types.Soldier {
	return &types.Soldier{
		Faction: faction,
		Health:  types.MaxHealth,

		Position: position,
		Velocity: types.Map2D{},
		Radius:   types.Map2D{X: 0.5, Y: 1},

		PointOfShooting: 0.6,
		ReloadingSpeed:  types.ReloadingSpeed,
		LastShot:        0,

		Actions: map[string]bool{},

		WeaponDamage: types.MaxDamage,
		Direction:    true,
	}

}
