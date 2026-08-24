package service

import (
	"spacetrash/game"
	"spacetrash/storage"
)

func (s *Session) Replay() (storage.Replay, error) {
	if s.Mission == nil {
		return storage.Replay{}, nil
	}
	return s.Store.LoadReplay(s.Mission.ID)
}

func (s *Session) MissionThreat() int {
	if s.Mission == nil {
		return 0
	}
	formation := game.BuildFormation(s.Mission.Asteroids)
	return game.FormationThreat(formation)
}
