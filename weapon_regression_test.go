package spacetrash

import (
	"path/filepath"
	"spacetrash/model"
	"spacetrash/service"
	"spacetrash/storage"
	"testing"
)

func TestWeaponUpgradeTouchesOnlySelectedSlot(t *testing.T) {
	store, err := storage.Open(filepath.Join(t.TempDir(), "regression.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	player := model.NewPlayer("p", "Nova")
	player.Level = 5
	player.Credits = 500
	vessel := model.NewVessel("v", player.ID, "Wayfarer")
	weapons := []model.Weapon{model.NewWeapon("w0", vessel.ID, "Pulse", 0), model.NewWeapon("w1", vessel.ID, "Arc", 1)}
	weapons[0].Level = 4
	weapons[1].Level = 4
	session, err := service.NewSession(store, player, vessel, weapons)
	if err != nil {
		t.Fatal(err)
	}
	beforeSelected := session.Weapons[0].Damage
	beforeAdjacent := session.Weapons[1].Damage
	if _, err := session.UpgradeWeapon(0); err != nil {
		t.Fatal(err)
	}
	if session.Weapons[0].Damage <= beforeSelected {
		t.Fatalf("selected weapon did not improve: %d", session.Weapons[0].Damage)
	}
	if session.Weapons[1].Damage != beforeAdjacent {
		t.Fatalf("adjacent weapon changed from %d to %d", beforeAdjacent, session.Weapons[1].Damage)
	}
}
