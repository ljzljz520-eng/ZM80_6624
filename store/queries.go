package store

import (
	"museumlogistics/model"
	"strings"
)

func (s *Store) SearchRecords(term string) ([]model.Record, error) {
	all, err := s.ListRecords()
	if err != nil {
		return nil, err
	}
	term = strings.ToLower(strings.TrimSpace(term))
	out := make([]model.Record, 0)
	for _, r := range all {
		if term == "" || strings.Contains(strings.ToLower(r.Title), term) || strings.Contains(strings.ToLower(r.Origin), term) || strings.Contains(strings.ToLower(r.Destination), term) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) RecordsByStatus(status string) ([]model.Record, error) {
	all, err := s.ListRecords()
	if err != nil {
		return nil, err
	}
	out := []model.Record{}
	for _, r := range all {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) SaveBundle(r model.Record, e model.Event, a model.Audit) error {
	if err := s.PutRecord(r); err != nil {
		return err
	}
	if err := s.PutEvent(e); err != nil {
		return err
	}
	return s.PutAudit(a)
}
