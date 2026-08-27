package intake

import (
	"context"
	"museumlogistics/store"
	"path/filepath"
	"testing"
)

func TestIntake(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r, e := (Service{Store: s}).Receive(context.Background(), "x", "Bowl", "A", "B")
	if e != nil || r.Status != "received" {
		t.Fatal(e, r)
	}
}
