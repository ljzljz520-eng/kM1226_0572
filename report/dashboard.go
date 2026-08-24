package report

import (
	"fmt"
	"spacetrash/game"
	"spacetrash/service"
	"strings"
)

func RenderDashboard(session *service.Session) string {
	telemetry := session.Telemetry()
	lines := []string{
		"SPACE DOCK TELEMETRY",
		fmt.Sprintf("pilot=%s level=%d lives=%d credits=%d", session.Player.Name, session.Player.Level, session.Player.Lives, session.Player.Credits),
		fmt.Sprintf("mission=%s hull=%d fuel=%d risk=%d alert=%s", session.MissionStatus(), session.Vessel.Hull, session.Vessel.Fuel, telemetry.Risk, telemetry.Alert),
		fmt.Sprintf("score=%d rank=%s return_to_dock=%t", telemetry.Score.Total, telemetry.Rank, telemetry.ReturnToDock),
		"plan=" + gamePlanLabel(telemetry.Plan),
	}
	return strings.Join(lines, "\n")
}

func gamePlanLabel(plan game.TacticalPlan) string {
	return game.OutcomeLabel(plan)
}

func RenderEvents(events []struct {
	Label    string
	Critical bool
}) string {
	lines := make([]string, 0, len(events))
	for _, event := range events {
		prefix := "-"
		if event.Critical {
			prefix = "!"
		}
		lines = append(lines, prefix+" "+event.Label)
	}
	return strings.Join(lines, "\n")
}

func RenderScore(score int, rank string) string {
	return fmt.Sprintf("Score %d (%s)", score, rank)
}

func RenderLoadout(session *service.Session) string {
	loadout := session.Loadout()
	lines := []string{fmt.Sprintf("loadout vessel=%s total_damage=%d energy=%d", loadout.VesselID, loadout.TotalDamage(), loadout.EnergyBudget())}
	for _, weapon := range loadout.Weapons {
		lines = append(lines, fmt.Sprintf("slot %d %s level=%d damage=%d", weapon.Slot, weapon.Name, weapon.Level, weapon.Damage))
	}
	return strings.Join(lines, "\n")
}

func RenderSurvival(session *service.Session) string {
	telemetry := session.Telemetry()
	return fmt.Sprintf("alert=%s risk=%d formation=%s dock=%t", telemetry.Alert, telemetry.Risk, session.FormationStatus(), telemetry.ReturnToDock)
}
