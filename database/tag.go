package database

import (
	"github.com/jmoiron/sqlx"
)

type Tag struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func GetTagByID(id int) (*Tag, error) {
	var tag Tag
	err := DB.Get(&tag, "SELECT * FROM tag WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func CreateTags(names []string) ([]*Tag, error) {
	for _, name := range names {
		_, err := DB.Exec(
			`INSERT INTO tag(name) VALUES ($1) ON CONFLICT DO NOTHING`,
			name,
		)
		if err != nil {
			return nil, err
		}
	}

	var tags []*Tag
	query, args, err := sqlx.In("SELECT * FROM tag WHERE name IN (?)", names)
	if err != nil {
		return nil, err
	}

	err = DB.Select(&tags, DB.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
