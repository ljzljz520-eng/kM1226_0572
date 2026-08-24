package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"go.etcd.io/bbolt"
	"spacetrash/game"
	"spacetrash/model"
)

var (
	playersBucket   = []byte("players")
	vesselsBucket   = []byte("vessels")
	weaponsBucket   = []byte("weapons")
	missionsBucket  = []byte("missions")
	asteroidsBucket = []byte("asteroids")
)

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, fmt.Errorf("database path is required")
	}
	db, err := bbolt.Open(path, 0o600, nil)
	if err != nil {
		return nil, err
	}
	store := &Store{db: db}
	if err := store.initialize(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return store, nil
}

func (s *Store) initialize() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range [][]byte{playersBucket, vesselsBucket, weaponsBucket, missionsBucket, asteroidsBucket} {
			if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}

func encode(value any) ([]byte, error) {
	return json.Marshal(value)
}

func decode(data []byte, target any) error {
	if len(data) == 0 {
		return fmt.Errorf("record is empty")
	}
	return json.Unmarshal(data, target)
}

func ensureOpen(s *Store) error {
	if s == nil || s.db == nil {
		return fmt.Errorf("store is closed")
	}
	return nil
}

func (s *Store) SavePlayer(player model.Player) error {
	if err := player.Validate(); err != nil {
		return err
	}
	data, err := encode(player)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(playersBucket).Put([]byte(player.ID), data) })
}

func (s *Store) SaveVessel(vessel model.Vessel) error {
	if err := vessel.Validate(); err != nil {
		return err
	}
	data, err := encode(vessel)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(vesselsBucket).Put([]byte(vessel.ID), data) })
}

func (s *Store) SaveWeapon(weapon model.Weapon) error {
	if err := weapon.Validate(); err != nil {
		return err
	}
	data, err := encode(weapon)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(weaponsBucket).Put([]byte(weapon.ID), data) })
}

func (s *Store) SaveMission(mission game.Mission) error {
	if err := mission.Validate(); err != nil {
		return err
	}
	data, err := encode(mission)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(missionsBucket).Put([]byte(mission.ID), data) })
}

func (s *Store) SaveAsteroid(asteroid model.Asteroid) error {
	if err := asteroid.Validate(); err != nil {
		return err
	}
	data, err := encode(asteroid)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(asteroidsBucket).Put([]byte(asteroid.ID), data) })
}

func SaveEntity[T any](s *Store, bucket, id string, value T) error {
	if id == "" || bucket == "" {
		return fmt.Errorf("entity key is required")
	}
	data, err := encode(value)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}
		return b.Put([]byte(id), data)
	})
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
