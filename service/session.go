package service

import (
	"fmt"
	"spacetrash/game"
	"spacetrash/model"
	"spacetrash/rules"
	"spacetrash/storage"
)

type Session struct {
	Store   *storage.Store
	Player  model.Player
	Vessel  model.Vessel
	Weapons []model.Weapon
	Mission *game.Mission
	Sector  game.Sector
}

func NewSession(store *storage.Store, player model.Player, vessel model.Vessel, weapons []model.Weapon) (*Session, error) {
	if store == nil {
		return nil, fmt.Errorf("store is required")
	}
	if err := rules.ValidateFleet(player, vessel, weapons); err != nil {
		return nil, err
	}
	copyWeapons := append([]model.Weapon(nil), weapons...)
	return &Session{Store: store, Player: player, Vessel: vessel, Weapons: copyWeapons, Sector: game.NewSector("Asterion")}, nil
}

func (s *Session) Save() error {
	if s.Mission == nil {
		return s.Store.SaveSnapshot(storage.Snapshot{Player: s.Player, Vessel: s.Vessel, Weapons: s.Weapons})
	}
	return s.Store.SaveSnapshot(storage.Snapshot{Player: s.Player, Vessel: s.Vessel, Weapons: s.Weapons, Missions: []game.Mission{*s.Mission}})
}

func (s *Session) BeginMission(id string, objective int, asteroids []model.Asteroid) error {
	if s.Mission != nil && s.Mission.Status == "active" {
		return fmt.Errorf("an active mission already exists")
	}
	mission := game.NewMission(id, s.Player.ID, s.Sector.Name, objective, asteroids)
	if err := mission.Validate(); err != nil {
		return err
	}
	s.Mission = &mission
	if _, err := s.RecordEvent("mission_started", "mission launched"); err != nil {
		return err
	}
	return nil
}

func (s *Session) Move(dx, dy int) error {
	if s.Mission == nil || s.Mission.Status != "active" {
		return fmt.Errorf("movement requires an active mission")
	}
	return game.Navigate(&s.Vessel, s.Sector, dx, dy)
}

func (s *Session) Scan(radius int) []model.Asteroid {
	if s.Mission == nil {
		return nil
	}
	return game.Scan(s.Vessel, s.Mission.Asteroids, radius)
}

func (s *Session) AdvanceTurn() error {
	if s.Mission == nil || s.Mission.Status != "active" {
		return fmt.Errorf("no active mission")
	}
	s.Mission.AdvanceTurn()
	if _, collided := game.ResolveCollision(&s.Vessel, s.Mission.Asteroids[0]); collided && !game.CanSurvive(s.Player, s.Vessel) {
		s.Mission.Status = "failed"
	}
	return nil
}
