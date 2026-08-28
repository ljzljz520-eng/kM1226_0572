package model

import "fmt"

type Asteroid struct {
	ID       string
	X        int
	Y        int
	Size     int
	Armor    int
	Reward   int
	Alive    bool
	Velocity int
}

func NewAsteroid(id string, x, y, size int) Asteroid {
	if size < 1 {
		size = 1
	}
	return Asteroid{ID: id, X: x, Y: y, Size: size, Armor: size * 6, Reward: size * 25, Alive: true, Velocity: size}
}

func (a Asteroid) Validate() error {
	if a.ID == "" || a.Size < 1 || a.Armor < 0 || a.Reward < 0 {
		return fmt.Errorf("asteroid data is invalid")
	}
	return nil
}

func (a *Asteroid) Drift() {
	a.X -= a.Velocity
	if a.X < -1000 {
		a.X = 1000
	}
}

func (a *Asteroid) Damage(amount int) bool {
	if !a.Alive || amount <= 0 {
		return false
	}
	a.Armor -= amount
	if a.Armor <= 0 {
		a.Alive = false
	}
	return !a.Alive
}

func (a Asteroid) ThreatDistance(x, y int) int {
	return abs(a.X-x) + abs(a.Y-y)
}

func (a Asteroid) Threatens(x, y, radius int) bool {
	return a.Alive && a.ThreatDistance(x, y) <= radius
}
