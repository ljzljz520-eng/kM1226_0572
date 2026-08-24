package storage

import (
	"path/filepath"
	"spacetrash/game"
	"spacetrash/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "space.db")
	store, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	player := model.NewPlayer("p1", "Nova")
	vessel := model.NewVessel("v1", player.ID, "Wayfarer")
	weapon := model.NewWeapon("w1", vessel.ID, "Pulse", 0)
	mission := game.NewMission("m1", player.ID, "Asterion", 2, []model.Asteroid{model.NewAsteroid("a1", 1, 1, 1)})
	for _, saveErr := range []error{store.SavePlayer(player), store.SaveVessel(vessel), store.SaveWeapon(weapon), store.SaveMission(mission)} {
		if saveErr != nil {
			t.Fatal(saveErr)
		}
	}
	if err := store.Close(); err != nil {
		t.Fatal(err)
	}
	store, err = Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	loaded, err := store.GetPlayer("p1")
	if err != nil || loaded.Name != "Nova" || loaded.Credits != 250 {
		t.Fatalf("player did not survive reopen: %+v %v", loaded, err)
	}
	loadedMission, err := store.GetMission("m1")
	if err != nil || loadedMission.Objective != 2 {
		t.Fatalf("mission did not survive reopen: %+v %v", loadedMission, err)
	}
}
