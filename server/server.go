package main

import (
	"log"
	"net/http"
	"os"
	"picollo"

	"picollo/model"

	"github.com/99designs/gqlgen/handler"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	session, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "picollo",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	repo := model.NewStorage(session)
	svc := model.NewService(repo)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(picollo.NewExecutableSchema(
		picollo.Config{
			Resolvers: &picollo.Resolver{Service: svc},
		})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
