package server

import (
	"museumlogistics/intake"
	"museumlogistics/query"
	"museumlogistics/store"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHTTP(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	h := (Server{Intake: intake.Service{Store: s}, Query: query.Service{Store: s}}).Handler()
	r := httptest.NewRequest("GET", "/records", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
