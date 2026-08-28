package rules

import (
	"fmt"
	"spacetrash/model"
)

func ValidateFleet(player model.Player, vessel model.Vessel, weapons []model.Weapon) error {
	if err := player.Validate(); err != nil {
		return err
	}
	if err := vessel.Validate(); err != nil {
		return err
	}
	if vessel.PlayerID != player.ID {
		return fmt.Errorf("vessel does not belong to player")
	}
	if len(weapons) == 0 {
		return fmt.Errorf("at least one weapon is required")
	}
	seen := make(map[int]bool)
	for _, weapon := range weapons {
		if err := weapon.Validate(); err != nil {
			return err
		}
		if weapon.VesselID != vessel.ID || seen[weapon.Slot] {
			return fmt.Errorf("weapon slots must be unique and belong to vessel")
		}
		seen[weapon.Slot] = true
	}
	return nil
}

func ClampCoordinate(value, limit int) int {
	if value < -limit {
		return -limit
	}
	if value > limit {
		return limit
	}
	return value
}

func MissionComplete(destroyed, required int) bool {
	if required <= 0 {
		return true
	}
	return destroyed >= required
}
