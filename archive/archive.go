package archive

import (
	"context"
	"fmt"
	"museumlogistics/model"
	"museumlogistics/store"
	"time"
)

type Service struct{ Store *store.Store }

func (s Service) Archive(ctx context.Context, id, actor string) (model.Audit, error) {
	if err := ctx.Err(); err != nil {
		return model.Audit{}, err
	}
	r, err := s.Store.GetRecord(id)
	if err != nil {
		return model.Audit{}, fmt.Errorf("archive load: %w", err)
	}
	if r.Status != "approved" {
		return model.Audit{}, model.ErrInvalid("archive status")
	}
	r.Transition("archived")
	a := model.Audit{ID: id + "-archive", RecordID: id, Action: "archive", Outcome: "success", At: now(), Reason: actor}
	e := model.Event{ID: id + "-archived", RecordID: id, Kind: "archived", Actor: actor, At: now()}
	return a, s.Store.SaveBundle(r, e, a)
}
func now() time.Time { return time.Now().UTC() }
