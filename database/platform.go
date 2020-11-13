package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Platform struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func GetAllPlatforms(db *sqlx.DB) ([]Platform, error) {
	var platforms []Platform
	err := db.Select(&platforms, "SELECT * FROM platform ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("Error getting platforms, %w", err)
	}

	return platforms, nil
}

func GetPlatformByID(db *sqlx.DB, id int) (*Platform, error) {
	var platform Platform
	err := db.Get(&platform, "SELECT * FROM platform WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting platform, %w", err)
	}
	return &platform, nil
}
