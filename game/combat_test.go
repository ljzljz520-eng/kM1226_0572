package game

import (
	"spacetrash/model"
	"testing"
)

func TestFireWeaponDestroysAsteroid(t *testing.T) {
	vessel := model.NewVessel("v", "p", "Wayfarer")
	weapon := model.NewWeapon("w", vessel.ID, "Pulse", 0)
	asteroid := model.NewAsteroid("a", 0, 0, 1)
	asteroid.Armor = 5
	result, err := FireWeapon(&vessel, &weapon, &asteroid)
	if err != nil || !result.Destroyed || result.Reward != asteroid.Reward {
		t.Fatalf("unexpected result: %+v %v", result, err)
	}
}

func TestSelectTarget(t *testing.T) {
	vessel := model.NewVessel("v", "p", "Wayfarer")
	target, ok := SelectTarget(vessel, []model.Asteroid{model.NewAsteroid("a", 50, 0, 1), model.NewAsteroid("b", 5, 0, 1)})
	if !ok || target.ID != "b" {
		t.Fatalf("unexpected target: %+v %t", target, ok)
	}
}
