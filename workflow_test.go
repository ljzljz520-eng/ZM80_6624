package main

import (
	"context"
	"museumlogistics/archive"
	"museumlogistics/intake"
	"museumlogistics/model"
	"museumlogistics/notify"
	"museumlogistics/query"
	"museumlogistics/review"
	"museumlogistics/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r, e := (intake.Service{Store: s}).Receive(context.Background(), "one", "Statue", "Paris", "Beijing")
	if e != nil || r.Status != "received" {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := model.NewRecord("two", "Mask", "Rome", "Beijing")
	s.PutRecord(r)
	r.Rule = "ok"
	s.PutRecord(r)
	review.Service{Store: s}.Approve(context.Background(), r.ID)
	a, e := (archive.Service{Store: s}).Archive(context.Background(), r.ID, "curator")
	if e != nil || a.Outcome != "success" {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	r := model.NewRecord("three", "Vase", "A", "B")
	r.Rule = "blocked"
	e := review.Evaluate(r)
	m, _ := notify.Build(context.Background(), r, e.Error())
	if m.Body == "" {
		t.Fatal("empty")
	}
	_ = query.Service{}
}
func TestBusinessChain38(t *testing.T) {
	r := model.NewRecord("bug", "Vase", "A", "B")
	r.Rule = "blocked"
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	s.PutRecord(r)
	_, err := (review.Service{Store: s}).Approve(context.Background(), r.ID)
	if err == nil || !review.IsRuleFailure(err) {
		t.Fatalf("审核失败无法追溯原始规则: %v", err)
	}
}
