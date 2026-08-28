package game

import "spacetrash/model"

type Formation struct {
	CenterX int
	CenterY int
	Spread  int
	Count   int
}

func BuildFormation(asteroids []model.Asteroid) Formation {
	if len(asteroids) == 0 {
		return Formation{}
	}
	minX, maxX := asteroids[0].X, asteroids[0].X
	minY, maxY := asteroids[0].Y, asteroids[0].Y
	for _, asteroid := range asteroids[1:] {
		if asteroid.X < minX {
			minX = asteroid.X
		}
		if asteroid.X > maxX {
			maxX = asteroid.X
		}
		if asteroid.Y < minY {
			minY = asteroid.Y
		}
		if asteroid.Y > maxY {
			maxY = asteroid.Y
		}
	}
	spread := maxX - minX
	if vertical := maxY - minY; vertical > spread {
		spread = vertical
	}
	return Formation{CenterX: (minX + maxX) / 2, CenterY: (minY + maxY) / 2, Spread: spread, Count: len(asteroids)}
}

func FormationThreat(formation Formation) int {
	if formation.Count == 0 {
		return 0
	}
	return formation.Count*formation.Spread + formation.Count*formation.Count
}

func IsDenseFormation(formation Formation) bool {
	return formation.Count >= 3 && formation.Spread <= formation.Count*10
}

func FormationTargets(vessel model.Vessel, asteroids []model.Asteroid) []model.Asteroid {
	formation := BuildFormation(asteroids)
	if formation.Count == 0 {
		return nil
	}
	if IsDenseFormation(formation) {
		return Scan(vessel, asteroids, formation.Spread+20)
	}
	return Scan(vessel, asteroids, 100)
}

func FormationCentroid(asteroids []model.Asteroid) (int, int) {
	if len(asteroids) == 0 {
		return 0, 0
	}
	x, y := 0, 0
	for _, asteroid := range asteroids {
		x += asteroid.X
		y += asteroid.Y
	}
	return x / len(asteroids), y / len(asteroids)
}

func FormationVelocity(asteroids []model.Asteroid) int {
	velocity := 0
	for _, asteroid := range asteroids {
		velocity += asteroid.Velocity
	}
	if len(asteroids) == 0 {
		return 0
	}
	return velocity / len(asteroids)
}

func FormationReward(asteroids []model.Asteroid) int {
	reward := 0
	for _, asteroid := range asteroids {
		if asteroid.Alive {
			reward += asteroid.Reward
		}
	}
	return reward
}

func FormationAlive(asteroids []model.Asteroid) int {
	count := 0
	for _, asteroid := range asteroids {
		if asteroid.Alive {
			count++
		}
	}
	return count
}

func FormationArmor(asteroids []model.Asteroid) int {
	armor := 0
	for _, asteroid := range asteroids {
		if asteroid.Alive {
			armor += asteroid.Armor
		}
	}
	return armor
}

func FormationStatus(asteroids []model.Asteroid) string {
	if len(asteroids) == 0 {
		return "empty"
	}
	if FormationAlive(asteroids) == 0 {
		return "cleared"
	}
	if IsDenseFormation(BuildFormation(asteroids)) {
		return "dense"
	}
	return "scattered"
}
