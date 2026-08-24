package model

import "fmt"

type Weapon struct {
	ID          string
	VesselID    string
	Slot        int
	Name        string
	Level       int
	Damage      int
	Cooldown    int
	EnergyCost  int
	UpgradeCost int
}

func NewWeapon(id, vesselID, name string, slot int) Weapon {
	return Weapon{ID: id, VesselID: vesselID, Slot: slot, Name: name, Level: 1, Damage: 12, Cooldown: 2, EnergyCost: 5, UpgradeCost: 40}
}

func (w Weapon) Validate() error {
	if w.ID == "" || w.VesselID == "" || w.Name == "" {
		return fmt.Errorf("weapon identity is required")
	}
	if w.Slot < 0 || w.Level < 1 || w.Damage <= 0 || w.Cooldown <= 0 || w.EnergyCost <= 0 {
		return fmt.Errorf("weapon statistics are invalid")
	}
	return nil
}

func (w Weapon) Upgradeable(playerLevel int) bool {
	return playerLevel >= w.Level && w.Level < 5
}

func (w *Weapon) Upgrade(playerLevel int) error {
	if !w.Upgradeable(playerLevel) {
		return fmt.Errorf("weapon cannot be upgraded at player level %d", playerLevel)
	}
	w.Level++
	w.Damage += 8
	if w.Cooldown > 1 {
		w.Cooldown--
	}
	w.EnergyCost++
	w.UpgradeCost += 20
	return nil
}

func (w Weapon) Fire(targetArmor int) (int, error) {
	if targetArmor < 0 {
		return 0, fmt.Errorf("target armor cannot be negative")
	}
	damage := w.Damage - targetArmor
	if damage < 1 {
		damage = 1
	}
	return damage, nil
}
