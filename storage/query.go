package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"spacetrash/game"
	"spacetrash/model"
)

func (s *Store) GetPlayer(id string) (model.Player, error) {
	var player model.Player
	err := s.read(playersBucket, id, &player)
	return player, err
}

func (s *Store) GetVessel(id string) (model.Vessel, error) {
	var vessel model.Vessel
	err := s.read(vesselsBucket, id, &vessel)
	return vessel, err
}

func (s *Store) GetWeapon(id string) (model.Weapon, error) {
	var weapon model.Weapon
	err := s.read(weaponsBucket, id, &weapon)
	return weapon, err
}

func (s *Store) GetMission(id string) (game.Mission, error) {
	var mission game.Mission
	err := s.read(missionsBucket, id, &mission)
	return mission, err
}

func (s *Store) GetAsteroid(id string) (model.Asteroid, error) {
	var asteroid model.Asteroid
	err := s.read(asteroidsBucket, id, &asteroid)
	return asteroid, err
}

func (s *Store) read(bucket []byte, id string, target any) error {
	if id == "" {
		return fmt.Errorf("record id is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("bucket does not exist")
		}
		data := b.Get([]byte(id))
		if data == nil {
			return fmt.Errorf("record %s not found", id)
		}
		return decode(data, target)
	})
}

func (s *Store) ListWeapons(vesselID string) ([]model.Weapon, error) {
	if vesselID == "" {
		return nil, fmt.Errorf("vessel id is required")
	}
	weapons := make([]model.Weapon, 0)
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return nil, err
	}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(weaponsBucket).ForEach(func(_, value []byte) error {
			var weapon model.Weapon
			if err := decode(value, &weapon); err != nil {
				return err
			}
			if weapon.VesselID == vesselID {
				weapons = append(weapons, weapon)
			}
			return nil
		})
	})
	return weapons, err
}

func (s *Store) ListMissions(playerID string) ([]game.Mission, error) {
	if playerID == "" {
		return nil, fmt.Errorf("player id is required")
	}
	missions := make([]game.Mission, 0)
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return nil, err
	}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(missionsBucket).ForEach(func(_, value []byte) error {
			var mission game.Mission
			if err := decode(value, &mission); err != nil {
				return err
			}
			if mission.PlayerID == playerID {
				missions = append(missions, mission)
			}
			return nil
		})
	})
	return missions, err
}

func (s *Store) Delete(bucket, id string) error {
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
		return b.Delete([]byte(id))
	})
}

func ReadBucketCount(s *Store, bucket string) (int, error) {
	count := 0
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return 0, err
	}
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}
		return b.ForEach(func(_, _ []byte) error { count++; return nil })
	})
	return count, err
}
