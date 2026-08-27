package query

import (
	"fmt"
	"museumlogistics/model"
	"strings"
)

func Render(records []model.Record) string {
	lines := make([]string, 0, len(records))
	for _, r := range records {
		lines = append(lines, r.Summary())
	}
	return fmt.Sprintf("%d records\n%s", len(records), strings.Join(lines, "\n"))
}
func Group(records []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range records {
		out[r.Status] = append(out[r.Status], r)
	}
	return out
}
func Latest(records []model.Record) (model.Record, bool) {
	if len(records) == 0 {
		return model.Record{}, false
	}
	best := records[0]
	for _, r := range records[1:] {
		if r.UpdatedAt.After(best.UpdatedAt) {
			best = r
		}
	}
	return best, true
}
