package game

import "spacetrash/model"

type SurvivalState struct {
	HullPercent int
	FuelPercent int
	Shield      int
	Lives       int
	CanContinue bool
	Alert       string
}

func EvaluateSurvival(player model.Player, vessel model.Vessel) SurvivalState {
	hull := 0
	if vessel.MaxHull > 0 {
		hull = vessel.Hull * 100 / vessel.MaxHull
	}
	fuel := 0
	if vessel.MaxFuel > 0 {
		fuel = vessel.Fuel * 100 / vessel.MaxFuel
	}
	alert := "nominal"
	if hull <= 25 || vessel.Shield == 0 {
		alert = "critical"
	} else if hull <= 50 || fuel <= 25 {
		alert = "caution"
	}
	return SurvivalState{HullPercent: hull, FuelPercent: fuel, Shield: vessel.Shield, Lives: player.Lives, CanContinue: CanSurvive(player, vessel), Alert: alert}
}

func ApplyEmergencyRepair(player *model.Player, vessel *model.Vessel) bool {
	if vessel.Hull > vessel.MaxHull/4 || player.Credits < 50 {
		return false
	}
	if err := player.SpendCredits(50); err != nil {
		return false
	}
	vessel.Repair(vessel.MaxHull / 2)
	return true
}

func ShouldReturnToDock(state SurvivalState, remaining int) bool {
	if !state.CanContinue {
		return true
	}
	if state.Alert == "critical" {
		return true
	}
	return state.FuelPercent < 10 && remaining > 0
}
