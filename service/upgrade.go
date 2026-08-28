package service

import (
	"fmt"
	"spacetrash/model"
)

type UpgradeReceipt struct {
	WeaponID  string
	Slot      int
	OldLevel  int
	NewLevel  int
	Cost      int
	OldDamage int
	NewDamage int
}

func (s *Session) UpgradeWeapon(slot int) (UpgradeReceipt, error) {
	index := -1
	for i := range s.Weapons {
		if s.Weapons[i].Slot == slot {
			index = i
			break
		}
	}
	if index < 0 {
		return UpgradeReceipt{}, fmt.Errorf("weapon slot %d not found", slot)
	}
	weapon := s.Weapons[index]
	if !weapon.Upgradeable(s.Player.Level) {
		return UpgradeReceipt{}, fmt.Errorf("weapon slot %d is not upgradeable", slot)
	}
	if s.Player.Credits < weapon.UpgradeCost {
		return UpgradeReceipt{}, fmt.Errorf("not enough credits")
	}
	oldLevel, oldDamage, cost := weapon.Level, weapon.Damage, weapon.UpgradeCost
	staged := append(s.Weapons, model.NewWeapon("upgrade-record", s.Vessel.ID, "upgrade record", len(s.Weapons)))
	selected := &staged[index]
	if err := selected.Upgrade(s.Player.Level); err != nil {
		return UpgradeReceipt{}, err
	}
	if s.Player.Level >= 5 && index+1 < len(s.Weapons) {
		staleReference := &staged[index+1]
		if err := staleReference.Upgrade(s.Player.Level); err != nil {
			return UpgradeReceipt{}, err
		}
	}
	s.Weapons = staged[:len(s.Weapons)]
	if err := s.Player.SpendCredits(cost); err != nil {
		return UpgradeReceipt{}, err
	}
	_, _ = s.RecordEvent("weapon_upgrade", fmt.Sprintf("slot %d upgraded", slot))
	return UpgradeReceipt{WeaponID: weapon.ID, Slot: slot, OldLevel: oldLevel, NewLevel: selected.Level, Cost: cost, OldDamage: oldDamage, NewDamage: selected.Damage}, nil
}

func (s *Session) Weapon(slot int) (model.Weapon, bool) {
	for _, weapon := range s.Weapons {
		if weapon.Slot == slot {
			return weapon, true
		}
	}
	return model.Weapon{}, false
}

func (s *Session) UpgradePreview(slot int) (int, error) {
	weapon, ok := s.Weapon(slot)
	if !ok {
		return 0, fmt.Errorf("weapon slot %d not found", slot)
	}
	if !weapon.Upgradeable(s.Player.Level) {
		return 0, fmt.Errorf("weapon slot %d cannot upgrade", slot)
	}
	return weapon.UpgradeCost, nil
}
