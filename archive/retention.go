package archive

import (
	"museumlogistics/model"
	"time"
)

func Eligible(r model.Record, age time.Duration) bool {
	return r.Status == "approved" && time.Since(r.UpdatedAt) >= age
}
func RetentionLabel(r model.Record) string {
	if r.Status == "archived" {
		return "permanent"
	}
	return "pending"
}
