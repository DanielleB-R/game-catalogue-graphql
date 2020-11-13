package schema

import (
	"github.com/DanielleB-R/game-catalogue-graphql/database"
	"github.com/graphql-go/graphql"
)

var PlatformType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Platform",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"games": &graphql.Field{
			Type: graphql.NewList(GameType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				platform := params.Source.(*database.Platform)
				return database.GetPlatformGames(platform.ID)
			},
		},
	},
})

var GameType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Game",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"platformID": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
