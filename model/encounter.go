package model

import "fmt"

type Encounter struct {
	ID        string
	MissionID string
	Name      string
	Threat    int
	Reward    int
	Resolved  bool
	Turn      int
}

func NewEncounter(id, missionID, name string, threat, reward int) Encounter {
	if threat < 1 {
		threat = 1
	}
	if reward < 0 {
		reward = 0
	}
	return Encounter{ID: id, MissionID: missionID, Name: name, Threat: threat, Reward: reward}
}

func (e Encounter) Validate() error {
	if e.ID == "" || e.MissionID == "" || e.Name == "" {
		return fmt.Errorf("encounter identity is required")
	}
	if e.Threat < 1 || e.Reward < 0 || e.Turn < 0 {
		return fmt.Errorf("encounter values are invalid")
	}
	return nil
}

func (e *Encounter) Resolve(playerLevel int) int {
	if e.Resolved || playerLevel < 1 {
		return 0
	}
	e.Resolved = true
	if playerLevel >= e.Threat {
		return e.Reward
	}
	return e.Reward / 2
}

func (e Encounter) Difficulty(playerLevel int) string {
	if playerLevel >= e.Threat {
		return "favorable"
	}
	if playerLevel+1 >= e.Threat {
		return "risky"
	}
	return "severe"
}

func (e *Encounter) Advance() {
	if !e.Resolved {
		e.Turn++
		if e.Turn > e.Threat*3 {
			e.Threat++
		}
	}
}

func (e Encounter) IsOverdue() bool {
	return !e.Resolved && e.Turn > e.Threat*3
}

func (e Encounter) Brief() string {
	status := "open"
	if e.Resolved {
		status = "resolved"
	}
	return fmt.Sprintf("%s (%s, threat %d, turn %d)", e.Name, status, e.Threat, e.Turn)
}
