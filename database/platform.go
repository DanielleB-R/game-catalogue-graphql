package database

import (
	"fmt"
)

type Platform struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func GetAllPlatforms() ([]Platform, error) {
	var platforms []Platform
	err := DB.Select(&platforms, "SELECT * FROM platform ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("Error getting platforms, %w", err)
	}

	return platforms, nil
}

func GetPlatformByID(id int) (*Platform, error) {
	var platform Platform
	err := DB.Get(&platform, "SELECT * FROM platform WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting platform, %w", err)
	}
	return &platform, nil
}

func CreatePlatform(name string) (*Platform, error) {
	newRow := DB.QueryRowx(
		"INSERT INTO platform(name) VALUES ($1) RETURNING *",
		name,
	)

	var platform Platform
	err := newRow.StructScan(&platform)
	if err != nil {
		return nil, fmt.Errorf("Error creating platform: %w", err)
	}
	return &platform, nil
}
