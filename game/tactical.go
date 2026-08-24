package game

import "spacetrash/model"

type TacticalPlan struct {
	Action      string
	DX          int
	DY          int
	Risk        int
	TargetID    string
	FuelNeeded  int
	Explanation string
}

func PlanTurn(vessel model.Vessel, asteroids []model.Asteroid, radius int) TacticalPlan {
	contacts := Scan(vessel, asteroids, radius)
	if len(contacts) == 0 {
		return TacticalPlan{Action: "hold", Explanation: "no contacts in scan radius"}
	}
	closest, _ := SelectTarget(vessel, contacts)
	distance := closest.ThreatDistance(vessel.X, vessel.Y)
	plan := TacticalPlan{Action: "fire", TargetID: closest.ID, Risk: closest.Size, Explanation: "closest live asteroid selected"}
	if distance <= 5 {
		plan.Action = "evade"
		plan.DX = 0
		plan.DY = 10
		plan.FuelNeeded = 10
		plan.Explanation = "collision range requires vertical evasion"
	}
	return plan
}

func EvasiveCourse(vessel model.Vessel, asteroid model.Asteroid) (int, int) {
	dx := 0
	dy := 10
	if asteroid.Y > vessel.Y {
		dy = -10
	}
	if vessel.Fuel < tacticalAbs(dy) {
		dx = -1
		dy = 0
	}
	return dx, dy
}

func RiskScore(vessel model.Vessel, asteroids []model.Asteroid) int {
	score := 0
	for _, asteroid := range asteroids {
		if asteroid.Threatens(vessel.X, vessel.Y, 25) {
			score += asteroid.Size * 10
		}
	}
	if vessel.Hull < vessel.MaxHull/2 {
		score += 20
	}
	if vessel.Fuel < vessel.MaxFuel/4 {
		score += 10
	}
	return score
}

func OutcomeLabel(plan TacticalPlan) string {
	if plan.Action == "evade" {
		return "evasive maneuver"
	}
	if plan.Action == "fire" {
		return "engage target " + plan.TargetID
	}
	return "hold position"
}

func tacticalAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
