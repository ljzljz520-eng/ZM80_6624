package notify

import (
	"context"
	"museumlogistics/model"
	"testing"
)

func TestNotify(t *testing.T) {
	m, e := Build(context.Background(), model.NewRecord("x", "v", "A", "B"), "ok")
	if e != nil || m.To != "A" {
		t.Fatal(e, m)
	}
	if Channel("approved") != "email" {
		t.Fatal("channel")
	}
}
