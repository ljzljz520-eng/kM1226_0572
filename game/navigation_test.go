package game

import (
	"spacetrash/model"
	"testing"
)

func TestNavigateAndScan(t *testing.T) {
	vessel := model.NewVessel("v", "p", "Wayfarer")
	if err := Navigate(&vessel, NewSector("A"), 5, 2); err != nil {
		t.Fatal(err)
	}
	asteroids := []model.Asteroid{model.NewAsteroid("a", 8, 2, 1), model.NewAsteroid("b", 200, 0, 1)}
	if got := len(Scan(vessel, asteroids, 10)); got != 1 {
		t.Fatalf("expected one contact, got %d", got)
	}
}

func TestNavigateRejectsInsufficientFuel(t *testing.T) {
	vessel := model.NewVessel("v", "p", "Wayfarer")
	vessel.Fuel = 1
	if err := Navigate(&vessel, NewSector("A"), 2, 0); err == nil {
		t.Fatal("expected fuel error")
	}
}
