package main

import (
	"net/http"

	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"

	"github.com/DanielleB-R/game-catalogue-graphql/schema"
)

func main() {
	graphqlSchema, err := schema.Get()
	if err != nil {
		panic(err)
	}

	handler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
