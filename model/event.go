package model

import (
	"fmt"
	"strings"
)

type Event struct {
	ID        string
	MissionID string
	Kind      string
	Message   string
	Turn      int
	Severity  string
}

func NewEvent(id, missionID, kind, message string, turn int) Event {
	severity := "info"
	if kind == "collision" || kind == "mission_failed" {
		severity = "warning"
	}
	if kind == "mission_complete" {
		severity = "success"
	}
	return Event{ID: id, MissionID: missionID, Kind: kind, Message: strings.TrimSpace(message), Turn: turn, Severity: severity}
}

func (e Event) Validate() error {
	if e.ID == "" || e.MissionID == "" || e.Kind == "" || e.Message == "" {
		return fmt.Errorf("event fields are required")
	}
	if e.Turn < 0 {
		return fmt.Errorf("event turn cannot be negative")
	}
	if e.Severity != "info" && e.Severity != "warning" && e.Severity != "success" {
		return fmt.Errorf("event severity is invalid")
	}
	return nil
}

func (e Event) IsCritical() bool {
	return e.Severity == "warning" || e.Kind == "mission_failed"
}

func (e Event) Label() string {
	return fmt.Sprintf("T%03d %s: %s", e.Turn, strings.ToUpper(e.Kind), e.Message)
}

func NormalizeEventKind(kind string) string {
	return strings.ToLower(strings.TrimSpace(kind))
}

func EventPriority(event Event) int {
	if event.IsCritical() {
		return 3
	}
	if event.Severity == "success" {
		return 2
	}
	return 1
}

func (e Event) IsMissionBoundary() bool {
	return e.Kind == "mission_started" || e.Kind == "mission_complete" || e.Kind == "mission_failed"
}

func EventSummary(events []Event) (int, int) {
	critical := 0
	boundaries := 0
	for _, event := range events {
		if event.IsCritical() {
			critical++
		}
		if event.IsMissionBoundary() {
			boundaries++
		}
	}
	return critical, boundaries
}
