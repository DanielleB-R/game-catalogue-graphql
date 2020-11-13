package database

import "github.com/jmoiron/sqlx"

type Game struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	PlatformID int    `db:"platform" json:"platformID"`
}

func GetAllGames(db *sqlx.DB) ([]Game, error) {
	var games []Game
	err := db.Select(&games, "SELECT * FROM game ORDER BY id")
	return games, err
}

func GetPlatformGames(db *sqlx.DB, platformID int) ([]Game, error) {
	var games []Game
	err := db.Select(&games, "SELECT * FROM game WHERE platform=$1 ORDER BY id", platformID)
	return games, err
}

func GetGameByID(db *sqlx.DB, id int) (*Game, error) {
	var game Game
	err := db.Get(&game, "SELECT * FROM game WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
