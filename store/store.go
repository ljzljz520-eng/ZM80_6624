package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"museumlogistics/model"
	"path/filepath"
	"sync"
)

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
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
func (s *Store) put(bucket, key string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		d := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if d == nil {
			return errors.New("not found")
		}
		return json.Unmarshal(d, v)
	})
}
func (s *Store) PutRecord(r model.Record) error { return s.put("records", r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	return r, s.get("records", id, &r)
}
func (s *Store) PutProfile(p model.Profile) error { return s.put("profiles", p.ID, p) }
func (s *Store) PutEvent(e model.Event) error     { return s.put("events", e.ID, e) }
func (s *Store) PutAudit(a model.Audit) error     { return s.put("audits", a.ID, a) }
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, err
}
