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
	},
})

var TagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tag",
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
				tag := params.Source.(*database.Tag)
				return database.GetGamesByTagID(tag.ID)
			},
		},
	},
})

func init() {
	GameType.AddFieldConfig("platform", &graphql.Field{
		Type: PlatformType,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			game := params.Source.(*database.Game)
			return database.GetPlatformByID(game.PlatformID)
		},
	})

	GameType.AddFieldConfig("tags", &graphql.Field{
		Type: graphql.NewList(TagType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			game := params.Source.(*database.Game)
			return database.GetTagsByGameID(game.ID)
		},
	})
}
