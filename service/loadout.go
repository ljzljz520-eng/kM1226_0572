package service

import "spacetrash/model"

func (s *Session) Loadout() model.Loadout {
	return model.NewLoadout(s.Vessel.ID, s.Weapons)
}

func (s *Session) SelectWeapon(slot int) error {
	loadout := s.Loadout()
	if err := loadout.Select(slot); err != nil {
		return err
	}
	return nil
}

func (s *Session) LoadoutDamage() int {
	return s.Loadout().TotalDamage()
}

func (s *Session) LoadoutEnergy() int {
	return s.Loadout().EnergyBudget()
}
