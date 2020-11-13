package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db := sqlx.MustConnect("postgres", os.Getenv("DB_URL"))

	var platformCount []int
	err := db.Select(&platformCount, "SELECT COUNT(*) FROM platform")
	if err != nil {
		panic(err)
	}

	fmt.Println("Platform count in database is: ", platformCount[0])
}
