package service

import (
	"spacetrash/game"
	"spacetrash/model"
	"spacetrash/rules"
)

type Telemetry struct {
	Risk         int
	Alert        string
	Score        rules.ScoreBreakdown
	Rank         string
	ReturnToDock bool
	Plan         game.TacticalPlan
}

func (s *Session) Telemetry() Telemetry {
	if s.Mission == nil {
		return Telemetry{Alert: "docked", ReturnToDock: true}
	}
	state := game.EvaluateSurvival(s.Player, s.Vessel)
	risk := game.RiskScore(s.Vessel, s.Mission.Asteroids)
	score := rules.MissionScore(s.Mission.Destroyed, s.Mission.Objective, s.Mission.Turn, s.Vessel.Hull, s.Vessel.MaxHull)
	return Telemetry{Risk: risk, Alert: state.Alert, Score: score, Rank: rules.RankForScore(score.Total), ReturnToDock: game.ShouldReturnToDock(state, s.Mission.Remaining()), Plan: game.PlanTurn(s.Vessel, s.Mission.Asteroids, 100)}
}

func (s *Session) ResolveEncounter(encounter *model.Encounter) (int, error) {
	if encounter == nil {
		return 0, nil
	}
	if err := encounter.Validate(); err != nil {
		return 0, err
	}
	reward := encounter.Resolve(s.Player.Level)
	if reward > 0 {
		s.Player.AwardCredits(reward)
		s.Player.AddXP(reward)
	}
	return reward, nil
}

func (s *Session) SurvivalAlert() string {
	return game.EvaluateSurvival(s.Player, s.Vessel).Alert
}

func (s *Session) ShouldDock() bool {
	if s.Mission == nil {
		return true
	}
	state := game.EvaluateSurvival(s.Player, s.Vessel)
	return game.ShouldReturnToDock(state, s.Mission.Remaining())
}

func (s *Session) FormationStatus() string {
	if s.Mission == nil {
		return "empty"
	}
	return game.FormationStatus(s.Mission.Asteroids)
}

func (s *Session) RemainingArmor() int {
	if s.Mission == nil {
		return 0
	}
	return game.FormationArmor(s.Mission.Asteroids)
}
