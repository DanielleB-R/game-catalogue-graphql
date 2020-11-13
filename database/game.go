package database

type Game struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	PlatformID int    `db:"platform" json:"platformID"`
}

func GetAllGames() ([]*Game, error) {
	var games []*Game
	err := DB.Select(&games, "SELECT * FROM game ORDER BY id")
	return games, err
}

func GetPlatformGames(platformID int) ([]*Game, error) {
	var games []*Game
	err := DB.Select(&games, "SELECT * FROM game WHERE platform=$1 ORDER BY id", platformID)
	return games, err
}

func GetGameByID(id int) (*Game, error) {
	var game Game
	err := DB.Get(&game, "SELECT * FROM game WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func GetGamesByTagID(tagID int) ([]*Game, error) {
	var games []*Game
	err := DB.Select(&games, "SELECT game.* FROM game INNER JOIN game_tag ON game.id = game_tag.game_id WHERE game_tag.tag_id=$1 ORDER BY id", tagID)
	return games, err
}

func CreateGame(name string, platformID int) (*Game, error) {
	newRow := DB.QueryRowx(
		"INSERT INTO game(name, platform) VALUES ($1, $2) RETURNING *",
		name,
		platformID,
	)

	var game Game
	err := newRow.StructScan(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
