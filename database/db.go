package database

import (
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	DB = sqlx.MustConnect("postgres", os.Getenv("DB_URL"))
}
