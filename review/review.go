package review

import (
	"context"
	"fmt"
	"museumlogistics/model"
	"museumlogistics/store"
)

type RuleError struct{ Rule, Reason string }

const MechanismID = "error.wrapping_lost"

func (e RuleError) Error() string { return fmt.Sprintf("rule %s failed: %s", e.Rule, e.Reason) }

type Service struct{ Store *store.Store }

func (s Service) Approve(ctx context.Context, id string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return model.Record{}, err
	}
	r, err := s.Store.GetRecord(id)
	if err != nil {
		return r, fmt.Errorf("load for review: %w", err)
	}
	if err := Evaluate(r); err != nil {
		return r, fmt.Errorf("review failed: %v", err)
	}
	if err := r.Transition("approved"); err != nil {
		return r, err
	}
	return r, s.Store.PutRecord(r)
}
func Evaluate(r model.Record) error {
	if r.Rule == "blocked" {
		return RuleError{Rule: "export-control", Reason: "origin requires curator approval"}
	}
	if r.Title == "" {
		return RuleError{Rule: "catalogue", Reason: "title missing"}
	}
	return nil
}
func Explain(err error) string {
	if err == nil {
		return "approved"
	}
	return "审核失败: " + err.Error()
}
