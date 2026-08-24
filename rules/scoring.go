package rules

import "spacetrash/model"

type ScoreBreakdown struct {
	DestroyPoints int
	SurvivalBonus int
	SpeedBonus    int
	Total         int
}

func MissionScore(destroyed, objective, turn, hull, maxHull int) ScoreBreakdown {
	if objective < 1 {
		objective = 1
	}
	if turn < 1 {
		turn = 1
	}
	destroyPoints := destroyed * 100
	survival := 0
	if maxHull > 0 {
		survival = hull * 50 / maxHull
	}
	speed := 0
	if destroyed >= objective {
		speed = objective * 20
		if turn <= objective*2 {
			speed += 50
		}
	}
	return ScoreBreakdown{DestroyPoints: destroyPoints, SurvivalBonus: survival, SpeedBonus: speed, Total: destroyPoints + survival + speed}
}

func RewardForAsteroid(asteroid model.Asteroid, chain int) int {
	if chain < 1 {
		chain = 1
	}
	reward := asteroid.Reward + chain*10
	if asteroid.Size >= 3 {
		reward += 50
	}
	return reward
}

func RankForScore(score int) string {
	if score >= 1000 {
		return "ace"
	}
	if score >= 500 {
		return "veteran"
	}
	if score >= 200 {
		return "pilot"
	}
	return "cadet"
}

func ValidateReward(amount int) bool {
	return amount >= 0 && amount <= 100000
}
