package types

import "github.com/graphql-go/graphql"

var SpeciesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SpeciesType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"classification": &graphql.Field{
				Type: graphql.String,
			},
			"designation": &graphql.Field{
				Type: graphql.String,
			},
			"averageHeight": &graphql.Field{
				Type: graphql.Int,
			},
			"skinColors": &graphql.Field{
				Type: graphql.String,
			},
			"hairColors": &graphql.Field{
				Type: graphql.String,
			},
			"eyeColors": &graphql.Field{
				Type: graphql.String,
			},
			"averageLifespan": &graphql.Field{
				Type: graphql.Int,
			},
			"language": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
