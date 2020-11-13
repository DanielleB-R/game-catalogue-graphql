package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/DanielleB-R/game-catalogue-graphql/schema"
)

func main() {
	db := sqlx.MustConnect("postgres", os.Getenv("DB_URL"))

	var platformCount []int
	err := db.Select(&platformCount, "SELECT COUNT(*) FROM platform")
	if err != nil {
		panic(err)
	}

	fmt.Println("Platform count in database is: ", platformCount[0])

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
