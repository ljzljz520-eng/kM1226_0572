package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

func AsJSON(summary MissionSummary) (string, error) {
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func RenderTable(summary MissionSummary) string {
	lines := []string{"SLOT | WEAPON | LEVEL | DAMAGE", "-----|--------|-------|-------"}
	for _, weapon := range summary.Weapons {
		lines = append(lines, fmt.Sprintf("%4d | %-12s | %5d | %6d", weapon.Slot, weapon.Name, weapon.Level, weapon.Damage))
	}
	return strings.Join(lines, "\n")
}

func RenderEvent(event string, details ...string) string {
	if len(details) == 0 {
		return "[EVENT] " + event
	}
	return "[EVENT] " + event + ": " + strings.Join(details, ", ")
}
