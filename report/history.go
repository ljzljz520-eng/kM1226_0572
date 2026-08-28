package report

import (
	"fmt"
	"spacetrash/model"
	"strings"
)

func RenderHistory(events []model.Event) string {
	lines := make([]string, 0, len(events))
	for _, event := range events {
		prefix := "INFO"
		if event.IsCritical() {
			prefix = "ALERT"
		}
		lines = append(lines, prefix+" "+event.Label())
	}
	return strings.Join(lines, "\n")
}

func HistoryStats(events []model.Event) (int, int, int) {
	critical := 0
	completed := 0
	maxTurn := 0
	for _, event := range events {
		if event.IsCritical() {
			critical++
		}
		if event.Kind == "mission_complete" {
			completed++
		}
		if event.Turn > maxTurn {
			maxTurn = event.Turn
		}
	}
	return critical, completed, maxTurn
}

func EventKinds(events []model.Event) []string {
	seen := make(map[string]bool)
	kinds := make([]string, 0)
	for _, event := range events {
		if !seen[event.Kind] {
			seen[event.Kind] = true
			kinds = append(kinds, event.Kind)
		}
	}
	return kinds
}

func RenderEventCount(events []model.Event) string {
	critical, completed, turns := HistoryStats(events)
	return fmt.Sprintf("events=%d critical=%d completed=%d turns=%d", len(events), critical, completed, turns)
}

func RenderKinds(events []model.Event) string {
	return strings.Join(EventKinds(events), ", ")
}

func CriticalEvents(events []model.Event) []model.Event {
	critical := make([]model.Event, 0)
	for _, event := range events {
		if event.IsCritical() {
			critical = append(critical, event)
		}
	}
	return critical
}
