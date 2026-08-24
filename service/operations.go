package service

import (
	"fmt"
	"spacetrash/game"
)

func (s *Session) Repair() error {
	cost := game.RepairCost(s.Vessel)
	if cost == 0 {
		return nil
	}
	if err := s.Player.SpendCredits(cost); err != nil {
		return err
	}
	s.Vessel.Repair(s.Vessel.MaxHull)
	return nil
}

func (s *Session) Refuel() error {
	cost := game.RefuelCost(s.Vessel)
	if cost == 0 {
		return nil
	}
	if err := s.Player.SpendCredits(cost); err != nil {
		return err
	}
	s.Vessel.Refuel()
	return nil
}

func (s *Session) RechargeShield(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("recharge amount must be positive")
	}
	if s.Player.Credits < amount {
		return fmt.Errorf("not enough credits")
	}
	s.Player.Credits -= amount
	s.Vessel.RechargeShield(amount)
	return nil
}

func (s *Session) MissionStatus() string {
	if s.Mission == nil {
		return "docked"
	}
	return s.Mission.Status
}
