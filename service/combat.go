package service

import (
	"fmt"
	"spacetrash/game"
	"spacetrash/model"
)

func (s *Session) Fire(slot int, asteroidID string) (game.AttackResult, error) {
	if s.Mission == nil || s.Mission.Status != "active" {
		return game.AttackResult{}, fmt.Errorf("fire requires an active mission")
	}
	weaponIndex := -1
	for i := range s.Weapons {
		if s.Weapons[i].Slot == slot {
			weaponIndex = i
			break
		}
	}
	if weaponIndex < 0 {
		return game.AttackResult{}, fmt.Errorf("weapon slot %d not found", slot)
	}
	asteroidIndex := -1
	for i := range s.Mission.Asteroids {
		if s.Mission.Asteroids[i].ID == asteroidID {
			asteroidIndex = i
			break
		}
	}
	if asteroidIndex < 0 {
		return game.AttackResult{}, fmt.Errorf("asteroid %s not found", asteroidID)
	}
	result, err := game.FireWeapon(&s.Vessel, &s.Weapons[weaponIndex], &s.Mission.Asteroids[asteroidIndex])
	if err != nil {
		return result, err
	}
	if result.Destroyed {
		s.Mission.RecordDestruction(result.Reward)
		game.ApplyReward(&s.Player, &s.Vessel, result.Reward, s.Mission.Destroyed)
		if s.Mission.Complete() {
			_, _ = s.RecordEvent("mission_complete", "debris objective cleared")
		} else {
			_, _ = s.RecordEvent("asteroid_destroyed", asteroidID)
		}
	} else {
		_, _ = s.RecordEvent("weapon_hit", asteroidID)
	}
	return result, nil
}

func (s *Session) Target() (model.Asteroid, bool) {
	if s.Mission == nil {
		return model.Asteroid{}, false
	}
	return game.SelectTarget(s.Vessel, s.Mission.Asteroids)
}

func (s *Session) DestroyedCount() int {
	if s.Mission == nil {
		return 0
	}
	return game.CountDestroyed(s.Mission.Asteroids)
}
