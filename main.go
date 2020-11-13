package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"

	"github.com/DanielleB-R/game-catalogue-graphql/database"
	"github.com/DanielleB-R/game-catalogue-graphql/schema"
)

func main() {
	platforms, err := database.GetAllPlatforms()
	if err != nil {
		panic(err)
	}

	fmt.Println("Platform count in database is: ", len(platforms))

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
