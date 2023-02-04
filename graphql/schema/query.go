package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/graphql/types"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"Vehicle": &graphql.Field{
				Type:        types.VehicleType,
				Description: "Get Vehicle by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// Resolve: ,
			},
			"Vehicles": &graphql.Field{
				Type:        graphql.NewList(types.VehicleType),
				Description: "Get Vehicle List",
				// Resolve: ,
			},
			"Starship": &graphql.Field{
				Type:        types.StarshipType,
				Description: "Get Starship by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// Resolve: ,
			},
			"Starships": &graphql.Field{
				Type:        graphql.NewList(types.StarshipType),
				Description: "Get Starship List",
				// Resolve: ,
			},
			"Species": &graphql.Field{
				Type:        types.SpeciesType,
				Description: "Get Species by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// Resolve: ,
			},
			"SpeciesList": &graphql.Field{
				Type:        graphql.NewList(types.SpeciesType),
				Description: "Get Species List",
				// Resolve: ,
			},
			"Planet": &graphql.Field{
				Type:        types.PlanetType,
				Description: "Get Planet by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// Resolve: ,
			},
			"Planets": &graphql.Field{
				Type:        graphql.NewList(types.PlanetType),
				Description: "Get Planet List",
				// Resolve: ,
			},
			"Character": &graphql.Field{
				Type:        types.CharacterType,
				Description: "Get Character by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// Resolve: ,
			},
			"Characters": &graphql.Field{
				Type:        graphql.NewList(types.CharacterType),
				Description: "Get Character List",
				// Resolve: ,
			},
		},
	},
)
