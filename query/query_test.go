package query

import (
	"museumlogistics/model"
	"testing"
)

func TestQuery(t *testing.T) {
	a := model.NewRecord("a", "Vase", "A", "B")
	b := model.NewRecord("b", "Mask", "C", "D")
	if len(Group([]model.Record{a, b})) != 1 {
		t.Fatal("group")
	}
	if _, ok := Latest(nil); ok {
		t.Fatal("latest")
	}
}
