package storage

import (
	"fmt"
	"spacetrash/model"
)

type Replay struct {
	MissionID string
	Events    []model.Event
}

func (s *Store) LoadReplay(missionID string) (Replay, error) {
	if missionID == "" {
		return Replay{}, fmt.Errorf("mission id is required")
	}
	events, err := s.ListEvents(missionID)
	if err != nil {
		return Replay{}, err
	}
	return Replay{MissionID: missionID, Events: events}, nil
}

func (r Replay) Duration() int {
	if len(r.Events) == 0 {
		return 0
	}
	return r.Events[len(r.Events)-1].Turn
}

func (r Replay) CriticalCount() int {
	count := 0
	for _, event := range r.Events {
		if event.IsCritical() {
			count++
		}
	}
	return count
}

func (r Replay) Labels() []string {
	labels := make([]string, 0, len(r.Events))
	for _, event := range r.Events {
		labels = append(labels, event.Label())
	}
	return labels
}

func (r Replay) HasKind(kind string) bool {
	for _, event := range r.Events {
		if event.Kind == kind {
			return true
		}
	}
	return false
}

func (r Replay) LastMessage() string {
	if len(r.Events) == 0 {
		return ""
	}
	return r.Events[len(r.Events)-1].Message
}

func (r Replay) Summary() string {
	return fmt.Sprintf("mission %s, %d events, %d critical, duration %d turns", r.MissionID, len(r.Events), r.CriticalCount(), r.Duration())
}

func (r Replay) FirstKind() string {
	if len(r.Events) == 0 {
		return ""
	}
	return r.Events[0].Kind
}

func (r Replay) EventAtTurn(turn int) []model.Event {
	selected := make([]model.Event, 0)
	for _, event := range r.Events {
		if event.Turn == turn {
			selected = append(selected, event)
		}
	}
	return selected
}

func (r Replay) Completed() bool {
	return r.HasKind("mission_complete")
}
