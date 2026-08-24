package model

import "testing"

func TestWeaponUpgrade(t *testing.T) {
	weapon := NewWeapon("w", "v", "Pulse", 0)
	if err := weapon.Upgrade(1); err != nil {
		t.Fatal(err)
	}
	if weapon.Level != 2 || weapon.Damage != 20 || weapon.Cooldown != 1 {
		t.Fatalf("unexpected weapon: %+v", weapon)
	}
}

func TestWeaponFire(t *testing.T) {
	weapon := NewWeapon("w", "v", "Pulse", 0)
	damage, err := weapon.Fire(100)
	if err != nil || damage != 1 {
		t.Fatalf("unexpected fire result %d %v", damage, err)
	}
}
