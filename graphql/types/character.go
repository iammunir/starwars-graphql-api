package types

import "github.com/graphql-go/graphql"

var CharacterType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CharacterType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"height": &graphql.Field{
				Type: graphql.Int,
			},
			"mass": &graphql.Field{
				Type: graphql.Int,
			},
			"hairColor": &graphql.Field{
				Type: graphql.String,
			},
			"skinColor": &graphql.Field{
				Type: graphql.String,
			},
			"eyeColor": &graphql.Field{
				Type: graphql.String,
			},
			"birthYear": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
