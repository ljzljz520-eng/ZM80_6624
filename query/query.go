package query

import (
	"context"
	"museumlogistics/model"
	"museumlogistics/store"
	"time"
)

type Service struct{ Store *store.Store }

func (s Service) Find(ctx context.Context, term string) ([]model.Record, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.Store.SearchRecords(term)
}
func (s Service) Timeline(ctx context.Context, id string) ([]model.Event, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return []model.Event{{ID: id + "-view", RecordID: id, Kind: "view", Actor: "operator", At: time.Now().UTC()}}, nil
}
