package review

import (
	"museumlogistics/model"
	"testing"
)

func TestReviewRule(t *testing.T) {
	r := model.NewRecord("x", "", "A", "B")
	if Evaluate(r) == nil {
		t.Fatal("expected rule error")
	}
	if !IsRuleFailure(Evaluate(r)) {
		t.Fatal("not rule failure")
	}
}
