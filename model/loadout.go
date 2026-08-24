package model

import "fmt"

type Loadout struct {
	VesselID string
	Weapons  []Weapon
	Active   int
}

func NewLoadout(vesselID string, weapons []Weapon) Loadout {
	copyWeapons := append([]Weapon(nil), weapons...)
	return Loadout{VesselID: vesselID, Weapons: copyWeapons, Active: 0}
}

func (l Loadout) Validate() error {
	if l.VesselID == "" || len(l.Weapons) == 0 {
		return fmt.Errorf("loadout requires a vessel and weapons")
	}
	if l.Active < 0 || l.Active >= len(l.Weapons) {
		return fmt.Errorf("active weapon index is invalid")
	}
	seen := make(map[int]bool)
	for _, weapon := range l.Weapons {
		if err := weapon.Validate(); err != nil {
			return err
		}
		if seen[weapon.Slot] {
			return fmt.Errorf("duplicate weapon slot")
		}
		seen[weapon.Slot] = true
	}
	return nil
}

func (l *Loadout) Select(slot int) error {
	for index, weapon := range l.Weapons {
		if weapon.Slot == slot {
			l.Active = index
			return nil
		}
	}
	return fmt.Errorf("slot %d is not installed", slot)
}

func (l Loadout) ActiveWeapon() Weapon {
	if l.Active < 0 || l.Active >= len(l.Weapons) {
		return Weapon{}
	}
	return l.Weapons[l.Active]
}

func (l Loadout) TotalDamage() int {
	total := 0
	for _, weapon := range l.Weapons {
		total += weapon.Damage
	}
	return total
}

func (l Loadout) EnergyBudget() int {
	budget := 0
	for _, weapon := range l.Weapons {
		budget += weapon.EnergyCost
	}
	return budget
}
