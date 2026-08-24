package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"sort"
	"spacetrash/model"
)

var eventsBucket = []byte("events")

func (s *Store) ensureEvents() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(eventsBucket)
		return err
	})
}

func (s *Store) SaveEvent(event model.Event) error {
	if err := event.Validate(); err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	if err := s.ensureEvents(); err != nil {
		return err
	}
	data, err := encode(event)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(eventsBucket).Put([]byte(event.ID), data) })
}

func (s *Store) ListEvents(missionID string) ([]model.Event, error) {
	if missionID == "" {
		return nil, fmt.Errorf("mission id is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return nil, err
	}
	events := make([]model.Event, 0)
	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(eventsBucket)
		if bucket == nil {
			return nil
		}
		return bucket.ForEach(func(_, data []byte) error {
			var event model.Event
			if err := decode(data, &event); err != nil {
				return err
			}
			if event.MissionID == missionID {
				events = append(events, event)
			}
			return nil
		})
	})
	sort.Slice(events, func(i, j int) bool { return events[i].Turn < events[j].Turn })
	return events, err
}

func (s *Store) DeleteEvents(missionID string) error {
	events, err := s.ListEvents(missionID)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(eventsBucket)
		if bucket == nil {
			return nil
		}
		for _, event := range events {
			if err := bucket.Delete([]byte(event.ID)); err != nil {
				return err
			}
		}
		return nil
	})
}
