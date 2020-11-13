package schema

import (
	"github.com/graphql-go/graphql"
)

func Get() (graphql.Schema, error) {
	platformFields := graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "Platform", Fields: platformFields}

	return graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
}
