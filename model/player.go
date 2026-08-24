package model

import "fmt"

type Player struct {
	ID       string
	Name     string
	Credits  int
	Lives    int
	Level    int
	XP       int
	VesselID string
}

func NewPlayer(id, name string) Player {
	return Player{ID: id, Name: name, Credits: 250, Lives: 3, Level: 1}
}

func (p Player) Validate() error {
	if p.ID == "" {
		return fmt.Errorf("player id is required")
	}
	if p.Name == "" {
		return fmt.Errorf("player name is required")
	}
	if p.Lives < 0 || p.Level < 1 || p.Credits < 0 || p.XP < 0 {
		return fmt.Errorf("player statistics are invalid")
	}
	return nil
}

func (p *Player) AwardCredits(amount int) error {
	if amount < 0 {
		return fmt.Errorf("credit award cannot be negative")
	}
	p.Credits += amount
	return nil
}

func (p *Player) SpendCredits(amount int) error {
	if amount < 0 || amount > p.Credits {
		return fmt.Errorf("insufficient credits")
	}
	p.Credits -= amount
	return nil
}

func (p *Player) AddXP(amount int) int {
	if amount > 0 {
		p.XP += amount
	}
	for p.XP >= p.Level*100 {
		p.XP -= p.Level * 100
		p.Level++
	}
	return p.Level
}

func (p *Player) LoseLife() bool {
	if p.Lives == 0 {
		return false
	}
	p.Lives--
	return true
}

func (p Player) Status() string {
	return fmt.Sprintf("%s level %d, lives %d, credits %d, xp %d", p.Name, p.Level, p.Lives, p.Credits, p.XP)
}
