package service

import "testing"

func TestRepairAndShield(t *testing.T) {
	session := testSession(t)
	session.Vessel.Hull = 60
	if err := session.Repair(); err != nil {
		t.Fatal(err)
	}
	if session.Vessel.Hull != session.Vessel.MaxHull {
		t.Fatalf("hull: %d", session.Vessel.Hull)
	}
	if err := session.RechargeShield(5); err != nil {
		t.Fatal(err)
	}
}

func TestUpgradePreview(t *testing.T) {
	session := testSession(t)
	cost, err := session.UpgradePreview(0)
	if err != nil || cost != session.Weapons[0].UpgradeCost {
		t.Fatalf("preview: %d %v", cost, err)
	}
}
