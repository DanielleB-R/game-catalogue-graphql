package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"

	"github.com/DanielleB-R/game-catalogue-graphql/database"
)

func Get(db *sqlx.DB) (graphql.Schema, error) {
	platformType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Platform",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
		"platform": &graphql.Field{
			Type:        platformType,
			Description: "Get a platform by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, isOK := params.Args["id"].(int)
				if isOK {
					platform, err := database.GetPlatformByID(db, id)
					if err == nil {
						return platform, nil
					}
				}
				return nil, nil
			},
		},
	}}

	return graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
}
