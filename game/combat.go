package game

import (
	"fmt"
	"spacetrash/model"
)

type AttackResult struct {
	WeaponID       string
	AsteroidID     string
	Damage         int
	Destroyed      bool
	Reward         int
	RemainingArmor int
}

func FireWeapon(vessel *model.Vessel, weapon *model.Weapon, asteroid *model.Asteroid) (AttackResult, error) {
	if vessel == nil || weapon == nil || asteroid == nil {
		return AttackResult{}, fmt.Errorf("combat participant is nil")
	}
	if !asteroid.Alive {
		return AttackResult{}, fmt.Errorf("asteroid is already destroyed")
	}
	if weapon.VesselID != vessel.ID {
		return AttackResult{}, fmt.Errorf("weapon belongs to another vessel")
	}
	damage, err := weapon.Fire(asteroid.Size)
	if err != nil {
		return AttackResult{}, err
	}
	destroyed := asteroid.Damage(damage)
	result := AttackResult{WeaponID: weapon.ID, AsteroidID: asteroid.ID, Damage: damage, Destroyed: destroyed, RemainingArmor: asteroid.Armor}
	if destroyed {
		result.Reward = asteroid.Reward
	}
	return result, nil
}

func ResolveCollision(vessel *model.Vessel, asteroid model.Asteroid) (int, bool) {
	if vessel == nil || !asteroid.Threatens(vessel.X, vessel.Y, 5) {
		return 0, false
	}
	damage := asteroid.Size * 5
	actual := vessel.Hit(damage)
	return actual, actual > 0
}

func SelectTarget(vessel model.Vessel, asteroids []model.Asteroid) (model.Asteroid, bool) {
	var selected model.Asteroid
	found := false
	for _, asteroid := range asteroids {
		if !asteroid.Alive {
			continue
		}
		if !asteroid.Threatens(vessel.X, vessel.Y, 100) {
			continue
		}
		if !found || asteroid.ThreatDistance(vessel.X, vessel.Y) < selected.ThreatDistance(vessel.X, vessel.Y) {
			selected = asteroid
			found = true
		}
	}
	return selected, found
}

func CountDestroyed(asteroids []model.Asteroid) int {
	count := 0
	for _, asteroid := range asteroids {
		if !asteroid.Alive {
			count++
		}
	}
	return count
}
