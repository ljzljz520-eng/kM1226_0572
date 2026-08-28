package service

import (
	"path/filepath"
	"spacetrash/model"
	"spacetrash/storage"
	"testing"
)

func testSession(t *testing.T) *Session {
	t.Helper()
	store, err := storage.Open(filepath.Join(t.TempDir(), "workflow.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { store.Close() })
	player := model.NewPlayer("p1", "Nova")
	vessel := model.NewVessel("v1", player.ID, "Wayfarer")
	weapons := []model.Weapon{model.NewWeapon("w1", vessel.ID, "Pulse", 0), model.NewWeapon("w2", vessel.ID, "Arc", 1)}
	session, err := NewSession(store, player, vessel, weapons)
	if err != nil {
		t.Fatal(err)
	}
	return session
}

func TestWorkflowOne(t *testing.T) {
	session := testSession(t)
	if err := session.BeginMission("m1", 1, []model.Asteroid{model.NewAsteroid("a1", 5, 0, 1)}); err != nil {
		t.Fatal(err)
	}
	if err := session.Move(2, 0); err != nil {
		t.Fatal(err)
	}
	if _, err := session.Fire(0, "a1"); err != nil {
		t.Fatal(err)
	}
	if session.MissionStatus() != "complete" {
		t.Fatalf("mission status: %s", session.MissionStatus())
	}
}

func TestWorkflowTwo(t *testing.T) {
	session := testSession(t)
	if err := session.BeginMission("m2", 2, []model.Asteroid{model.NewAsteroid("a1", 2, 0, 1), model.NewAsteroid("a2", 3, 0, 1)}); err != nil {
		t.Fatal(err)
	}
	if _, ok := session.Target(); !ok {
		t.Fatal("expected target")
	}
	if _, err := session.Fire(0, "a1"); err != nil {
		t.Fatal(err)
	}
	if session.DestroyedCount() != 1 {
		t.Fatalf("destroyed count: %d", session.DestroyedCount())
	}
}

func TestWorkflowThree(t *testing.T) {
	session := testSession(t)
	session.Vessel.Fuel = 10
	if err := session.BeginMission("m3", 1, []model.Asteroid{model.NewAsteroid("a1", 50, 0, 1)}); err != nil {
		t.Fatal(err)
	}
	if err := session.Move(3, 0); err != nil {
		t.Fatal(err)
	}
	if err := session.Refuel(); err != nil {
		t.Fatal(err)
	}
	if session.Vessel.Fuel != session.Vessel.MaxFuel {
		t.Fatalf("fuel was not restored: %d", session.Vessel.Fuel)
	}
	if err := session.Save(); err != nil {
		t.Fatal(err)
	}
}
