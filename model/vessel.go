package model

import "fmt"

type Vessel struct {
	ID       string
	PlayerID string
	Name     string
	X        int
	Y        int
	Hull     int
	MaxHull  int
	Fuel     int
	MaxFuel  int
	Shield   int
}

func NewVessel(id, playerID, name string) Vessel {
	return Vessel{ID: id, PlayerID: playerID, Name: name, MaxHull: 100, Hull: 100, MaxFuel: 100, Fuel: 100, Shield: 25}
}

func (v Vessel) Validate() error {
	if v.ID == "" || v.PlayerID == "" {
		return fmt.Errorf("vessel identity is required")
	}
	if v.MaxHull <= 0 || v.Hull < 0 || v.Hull > v.MaxHull || v.Fuel < 0 || v.Fuel > v.MaxFuel {
		return fmt.Errorf("vessel resources are invalid")
	}
	return nil
}

func (v *Vessel) Move(dx, dy int) error {
	if dx == 0 && dy == 0 {
		return fmt.Errorf("movement must change position")
	}
	cost := abs(dx) + abs(dy)
	if cost > v.Fuel {
		return fmt.Errorf("not enough fuel")
	}
	v.X += dx
	v.Y += dy
	v.Fuel -= cost
	return nil
}

func (v *Vessel) Refuel() int {
	added := v.MaxFuel - v.Fuel
	v.Fuel = v.MaxFuel
	return added
}

func (v *Vessel) Repair(amount int) int {
	if amount < 0 {
		return 0
	}
	before := v.Hull
	v.Hull += amount
	if v.Hull > v.MaxHull {
		v.Hull = v.MaxHull
	}
	return v.Hull - before
}

func (v *Vessel) Hit(amount int) int {
	if amount <= 0 {
		return 0
	}
	absorbed := amount
	if v.Shield >= absorbed {
		v.Shield -= absorbed
		return 0
	}
	absorbed -= v.Shield
	v.Shield = 0
	if absorbed > v.Hull {
		absorbed = v.Hull
	}
	v.Hull -= absorbed
	return absorbed
}

func (v *Vessel) RechargeShield(amount int) int {
	if amount <= 0 {
		return 0
	}
	v.Shield += amount
	if v.Shield > 100 {
		v.Shield = 100
	}
	return v.Shield
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
