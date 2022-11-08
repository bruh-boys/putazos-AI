package models

import (
	"math/rand"

	"github.com/bruh-boys/putazos-ai/prototype/types"
)

func NewSoldier(faction string, position types.Map2D) *types.Soldier {
	return &types.Soldier{
		Faction: faction,
		Health:  types.MaxHealth,

		Position: position,
		Velocity: types.Map2D{},
		Radius:   types.Map2D{X: 16, Y: 38},

		PointOfShooting: 0.6,
		ReloadingSpeed:  types.ReloadingSpeed,
		LastShot:        0,

		Actions: map[string]bool{},

		WeaponDamage: types.MaxDamage,
		Direction:    true,
		Id:           string(rune(rand.Intn(100000))),
	}

}
