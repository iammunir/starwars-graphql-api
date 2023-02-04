package types

import "github.com/graphql-go/graphql"

var VehicleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VehicleType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"model": &graphql.Field{
				Type: graphql.String,
			},
			"manufacturer": &graphql.Field{
				Type: graphql.String,
			},
			"costInCredits": &graphql.Field{
				Type: graphql.Int,
			},
			"length": &graphql.Field{
				Type: graphql.Float,
			},
			"maxAtmospheringSpeed": &graphql.Field{
				Type: graphql.Int,
			},
			"crew": &graphql.Field{
				Type: graphql.Int,
			},
			"passengers": &graphql.Field{
				Type: graphql.Int,
			},
			"cargoCapacity": &graphql.Field{
				Type: graphql.Int,
			},
			"consumables": &graphql.Field{
				Type: graphql.String,
			},
			"vehicleClass": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
