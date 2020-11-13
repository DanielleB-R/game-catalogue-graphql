package schema

import "github.com/graphql-go/graphql"

var PlatformType = graphql.NewObject(graphql.ObjectConfig{
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
