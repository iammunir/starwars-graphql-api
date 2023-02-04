package types

import "github.com/graphql-go/graphql"

var PlanetType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlanetType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"rotationPeriod": &graphql.Field{
				Type: graphql.Int,
			},
			"orbitalPeriod": &graphql.Field{
				Type: graphql.Int,
			},
			"diameter": &graphql.Field{
				Type: graphql.Int,
			},
			"climate": &graphql.Field{
				Type: graphql.String,
			},
			"gravity": &graphql.Field{
				Type: graphql.String,
			},
			"terrain": &graphql.Field{
				Type: graphql.String,
			},
			"surfaceWater": &graphql.Field{
				Type: graphql.Int,
			},
			"population": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
