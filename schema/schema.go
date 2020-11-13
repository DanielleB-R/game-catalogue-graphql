package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/DanielleB-R/game-catalogue-graphql/database"
)

func Get() (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
		"platform": &graphql.Field{
			Type:        PlatformType,
			Description: "Get a platform by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, isOK := params.Args["id"].(int)
				if isOK {
					platform, err := database.GetPlatformByID(id)
					if err == nil {
						return platform, nil
					}
				}
				return nil, nil
			},
		},
		"platforms": &graphql.Field{
			Type:        graphql.NewList(PlatformType),
			Description: "All of the platforms",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return database.GetAllPlatforms()
			},
		},
		"game": &graphql.Field{
			Type:        GameType,
			Description: "Get a game by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, isOK := params.Args["id"].(int)
				if isOK {
					game, err := database.GetGameByID(id)
					if err == nil {
						return game, nil
					}
				}
				return nil, nil
			},
		},
	}})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPlatform": &graphql.Field{
				Type:        PlatformType,
				Description: "Add a new gaming platform",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)

					return database.CreatePlatform(name)
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
