package intake

import (
	"context"
	"fmt"
	"museumlogistics/model"
	"museumlogistics/store"
)

type Service struct{ Store *store.Store }

func (s Service) Register(ctx context.Context, r model.Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := r.Validate(); err != nil {
		return err
	}
	if r.Status == "" {
		r.Status = "received"
	}
	return s.Store.PutRecord(r)
}
func (s Service) Receive(ctx context.Context, id, title, origin, destination string) (model.Record, error) {
	r := model.NewRecord(id, title, origin, destination)
	if err := s.Register(ctx, r); err != nil {
		return r, fmt.Errorf("receive %s: %w", id, err)
	}
	return r, nil
}
func (s Service) AddProfile(ctx context.Context, p model.Profile) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if p.ID == "" || p.Institution == "" {
		return model.ErrInvalid("profile")
	}
	return s.Store.PutProfile(p)
}
