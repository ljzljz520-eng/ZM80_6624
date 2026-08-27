package archive

import (
	"museumlogistics/model"
	"testing"
	"time"
)

func TestRetention(t *testing.T) {
	r := model.NewRecord("x", "v", "A", "B")
	if Eligible(r, time.Hour) {
		t.Fatal("not eligible")
	}
	r.Status = "archived"
	if RetentionLabel(r) != "permanent" {
		t.Fatal("label")
	}
}
