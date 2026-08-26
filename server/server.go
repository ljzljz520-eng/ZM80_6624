package server

import (
	"context"
	"encoding/json"
	"museumlogistics/intake"
	"museumlogistics/query"
	"net/http"
)

type Server struct {
	Intake intake.Service
	Query  query.Service
}

func (s Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/records", s.records)
	return mux
}
func (s Server) records(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method == http.MethodPost {
		var in struct{ ID, Title, Origin, Destination string }
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			http.Error(w, "invalid json", 400)
			return
		}
		rec, err := s.Intake.Receive(ctx, in.ID, in.Title, in.Origin, in.Destination)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		json.NewEncoder(w).Encode(rec)
		return
	}
	items, err := s.Query.Find(ctx, r.URL.Query().Get("q"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(items)
}
func Run(ctx context.Context, addr string, h http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: h}
	go func() { <-ctx.Done(); srv.Shutdown(context.Background()) }()
	return srv.ListenAndServe()
}
