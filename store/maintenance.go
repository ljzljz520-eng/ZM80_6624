package store

import (
	"errors"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Health() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		if tx.Bucket([]byte("records")) == nil {
			return errors.New("records bucket missing")
		}
		return nil
	})
}
func (s *Store) PurgeBefore(cutoff time.Time) (int, error) {
	all, err := s.ListRecords()
	if err != nil {
		return 0, err
	}
	n := 0
	for _, r := range all {
		if r.UpdatedAt.Before(cutoff) {
			if e := s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(r.ID)) }); e != nil {
				return n, e
			}
			n++
		}
	}
	return n, nil
}
