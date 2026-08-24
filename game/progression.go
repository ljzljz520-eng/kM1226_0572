package game

import "spacetrash/model"

type Progression struct {
	LevelBefore int
	LevelAfter  int
	Credits     int
	Lives       int
}

func ApplyReward(player *model.Player, vessel *model.Vessel, reward, salvage int) Progression {
	before := player.Level
	if reward > 0 {
		_ = player.AwardCredits(reward)
	}
	if salvage > 0 {
		_ = player.AwardCredits(salvage * 5)
	}
	player.AddXP(reward + salvage*10)
	if vessel.Hull == 0 && player.LoseLife() {
		vessel.Hull = vessel.MaxHull / 2
	}
	return Progression{LevelBefore: before, LevelAfter: player.Level, Credits: player.Credits, Lives: player.Lives}
}

func RepairCost(vessel model.Vessel) int {
	missing := vessel.MaxHull - vessel.Hull
	if missing <= 0 {
		return 0
	}
	return missing * 2
}

func RefuelCost(vessel model.Vessel) int {
	missing := vessel.MaxFuel - vessel.Fuel
	if missing <= 0 {
		return 0
	}
	return missing
}

func CanSurvive(player model.Player, vessel model.Vessel) bool {
	return player.Lives > 0 && vessel.Hull > 0
}
