package schema

import "github.com/graphql-go/graphql"

func NewGraphQLSchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    RootQuery,
			Mutation: nil,
		},
	)
}
