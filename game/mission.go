package game

import (
	"fmt"
	"spacetrash/model"
)

type Mission struct {
	ID        string
	PlayerID  string
	Sector    string
	Objective int
	Destroyed int
	Reward    int
	Status    string
	Turn      int
	Asteroids []model.Asteroid
}

func NewMission(id, playerID, sector string, objective int, asteroids []model.Asteroid) Mission {
	if objective < 1 {
		objective = 1
	}
	return Mission{ID: id, PlayerID: playerID, Sector: sector, Objective: objective, Reward: objective * 75, Status: "active", Asteroids: asteroids}
}

func (m *Mission) Validate() error {
	if m.ID == "" || m.PlayerID == "" || m.Sector == "" {
		return fmt.Errorf("mission identity is required")
	}
	if m.Objective < 1 || m.Destroyed < 0 || m.Destroyed > m.Objective {
		return fmt.Errorf("mission progress is invalid")
	}
	if m.Status != "active" && m.Status != "complete" && m.Status != "failed" {
		return fmt.Errorf("mission status is invalid")
	}
	return nil
}

func (m *Mission) RecordDestruction(reward int) bool {
	if m.Status != "active" {
		return false
	}
	if reward > 0 {
		m.Reward += reward
	}
	m.Destroyed++
	if m.Destroyed >= m.Objective {
		m.Status = "complete"
	}
	return true
}

func (m *Mission) AdvanceTurn() {
	if m.Status == "active" {
		m.Turn++
		for i := range m.Asteroids {
			m.Asteroids[i].Drift()
		}
	}
}

func (m Mission) Remaining() int {
	remaining := m.Objective - m.Destroyed
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (m Mission) Complete() bool {
	return m.Status == "complete"
}
