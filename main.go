package main

import (
	"context"
	"log"
	"museumlogistics/intake"
	"museumlogistics/query"
	"museumlogistics/server"
	"museumlogistics/store"
)

func main() {
	st, err := store.Open("museum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
	svc := server.Server{Intake: intake.Service{Store: st}, Query: query.Service{Store: st}}
	if err := server.Run(context.Background(), ":8080", svc.Handler()); err != nil {
		log.Print(err)
	}
}
