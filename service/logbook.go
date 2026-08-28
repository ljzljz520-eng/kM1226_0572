package service

import (
	"fmt"
	"spacetrash/model"
)

func (s *Session) RecordEvent(kind, message string) (model.Event, error) {
	if s.Mission == nil {
		return model.Event{}, fmt.Errorf("event requires a mission")
	}
	eventID := fmt.Sprintf("%s-%d-%s", s.Mission.ID, s.Mission.Turn, kind)
	event := model.NewEvent(eventID, s.Mission.ID, model.NormalizeEventKind(kind), message, s.Mission.Turn)
	if err := s.Store.SaveEvent(event); err != nil {
		return model.Event{}, err
	}
	return event, nil
}

func (s *Session) Events() ([]model.Event, error) {
	if s.Mission == nil {
		return nil, fmt.Errorf("event history requires a mission")
	}
	return s.Store.ListEvents(s.Mission.ID)
}

func (s *Session) EventCount() int {
	events, err := s.Events()
	if err != nil {
		return 0
	}
	return len(events)
}

func (s *Session) SaveEventAndSnapshot(kind, message string) error {
	if _, err := s.RecordEvent(kind, message); err != nil {
		return err
	}
	return s.Save()
}
