package game

import (
	"fmt"
	"spacetrash/model"
	"spacetrash/rules"
)

type Sector struct {
	Name      string
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	SafeRange int
}

func NewSector(name string) Sector {
	return Sector{Name: name, MinX: -500, MaxX: 500, MinY: -500, MaxY: 500, SafeRange: 30}
}

func Navigate(vessel *model.Vessel, sector Sector, dx, dy int) error {
	if vessel == nil {
		return fmt.Errorf("vessel is nil")
	}
	if dx == 0 && dy == 0 {
		return fmt.Errorf("course is empty")
	}
	if vessel.X+dx < sector.MinX || vessel.X+dx > sector.MaxX || vessel.Y+dy < sector.MinY || vessel.Y+dy > sector.MaxY {
		return fmt.Errorf("course leaves sector")
	}
	if err := vessel.Move(dx, dy); err != nil {
		return err
	}
	vessel.X = rules.ClampCoordinate(vessel.X, 1000)
	vessel.Y = rules.ClampCoordinate(vessel.Y, 1000)
	return nil
}

func Scan(vessel model.Vessel, asteroids []model.Asteroid, radius int) []model.Asteroid {
	visible := make([]model.Asteroid, 0, len(asteroids))
	for _, asteroid := range asteroids {
		if asteroid.Threatens(vessel.X, vessel.Y, radius) {
			visible = append(visible, asteroid)
		}
	}
	return visible
}

func DriftField(asteroids []model.Asteroid) []model.Asteroid {
	updated := make([]model.Asteroid, len(asteroids))
	copy(updated, asteroids)
	for i := range updated {
		updated[i].Drift()
	}
	return updated
}
