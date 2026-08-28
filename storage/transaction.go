package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"spacetrash/game"
	"spacetrash/model"
)

type Snapshot struct {
	Player   model.Player
	Vessel   model.Vessel
	Weapons  []model.Weapon
	Missions []game.Mission
}

func (s *Store) SaveSnapshot(snapshot Snapshot) error {
	if err := snapshot.Player.Validate(); err != nil {
		return err
	}
	if err := snapshot.Vessel.Validate(); err != nil {
		return err
	}
	if len(snapshot.Weapons) == 0 {
		return fmt.Errorf("snapshot has no weapons")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ensureOpen(s); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := putJSON(tx.Bucket(playersBucket), snapshot.Player.ID, snapshot.Player); err != nil {
			return err
		}
		if err := putJSON(tx.Bucket(vesselsBucket), snapshot.Vessel.ID, snapshot.Vessel); err != nil {
			return err
		}
		for _, weapon := range snapshot.Weapons {
			if err := putJSON(tx.Bucket(weaponsBucket), weapon.ID, weapon); err != nil {
				return err
			}
		}
		for _, mission := range snapshot.Missions {
			if err := putJSON(tx.Bucket(missionsBucket), mission.ID, mission); err != nil {
				return err
			}
		}
		return nil
	})
}

func putJSON(bucket *bbolt.Bucket, id string, value any) error {
	if bucket == nil || id == "" {
		return fmt.Errorf("invalid snapshot record")
	}
	data, err := encode(value)
	if err != nil {
		return err
	}
	return bucket.Put([]byte(id), data)
}

func (s *Store) LoadSnapshot(playerID, vesselID string) (Snapshot, error) {
	player, err := s.GetPlayer(playerID)
	if err != nil {
		return Snapshot{}, err
	}
	vessel, err := s.GetVessel(vesselID)
	if err != nil {
		return Snapshot{}, err
	}
	weapons, err := s.ListWeapons(vesselID)
	if err != nil {
		return Snapshot{}, err
	}
	missions, err := s.ListMissions(playerID)
	if err != nil {
		return Snapshot{}, err
	}
	return Snapshot{Player: player, Vessel: vessel, Weapons: weapons, Missions: missions}, nil
}
