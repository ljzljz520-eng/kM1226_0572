package report

import (
	"fmt"
	"sort"
	"spacetrash/service"
)

type MissionSummary struct {
	PlayerName string
	Level      int
	Lives      int
	Credits    int
	Position   string
	Hull       int
	Fuel       int
	Mission    string
	Remaining  int
	Weapons    []WeaponSummary
}

type WeaponSummary struct {
	Slot   int
	Name   string
	Level  int
	Damage int
}

func BuildSummary(session *service.Session) MissionSummary {
	result := MissionSummary{PlayerName: session.Player.Name, Level: session.Player.Level, Lives: session.Player.Lives, Credits: session.Player.Credits, Position: fmt.Sprintf("%d,%d", session.Vessel.X, session.Vessel.Y), Hull: session.Vessel.Hull, Fuel: session.Vessel.Fuel, Mission: session.MissionStatus()}
	if session.Mission != nil {
		result.Remaining = session.Mission.Remaining()
	}
	for _, weapon := range session.Weapons {
		result.Weapons = append(result.Weapons, WeaponSummary{Slot: weapon.Slot, Name: weapon.Name, Level: weapon.Level, Damage: weapon.Damage})
	}
	sort.Slice(result.Weapons, func(i, j int) bool { return result.Weapons[i].Slot < result.Weapons[j].Slot })
	return result
}

func RenderSummary(summary MissionSummary) string {
	text := fmt.Sprintf("Pilot %s | level %d | lives %d | credits %d\n", summary.PlayerName, summary.Level, summary.Lives, summary.Credits)
	text += fmt.Sprintf("Position %s | hull %d | fuel %d | mission %s | remaining %d\n", summary.Position, summary.Hull, summary.Fuel, summary.Mission, summary.Remaining)
	for _, weapon := range summary.Weapons {
		text += fmt.Sprintf("Slot %d: %s level %d damage %d\n", weapon.Slot, weapon.Name, weapon.Level, weapon.Damage)
	}
	return text
}

func RenderMission(session *service.Session) string {
	if session.Mission == nil {
		return "No active mission"
	}
	return fmt.Sprintf("Mission %s in %s: %s, destroyed %d/%d, turn %d", session.Mission.ID, session.Mission.Sector, session.Mission.Status, session.Mission.Destroyed, session.Mission.Objective, session.Mission.Turn)
}
