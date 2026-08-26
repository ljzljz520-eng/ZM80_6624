package intake

import (
	"context"
	"museumlogistics/model"
)

func BuildManifest(ctx context.Context, r model.Record, p model.Profile) (map[string]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if !p.Approved {
		return nil, model.ErrInvalid("profile approval")
	}
	return map[string]string{"record": r.ID, "institution": p.Institution, "route": r.Origin + "->" + r.Destination}, nil
}
func NormalizeRecord(r model.Record) model.Record {
	if r.Origin == "" {
		r.Origin = "unknown"
	}
	if r.Destination == "" {
		r.Destination = "unknown"
	}
	return r
}
