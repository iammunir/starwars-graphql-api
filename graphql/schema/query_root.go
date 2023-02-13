package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/graphql/resolver"
	"github.com/iammunir/starwars-graphql-api/graphql/types"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"vehicle": &graphql.Field{
				Type:        types.VehicleType,
				Description: "Get Vehicle by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.VehicleResolver,
			},
			"vehicles": &graphql.Field{
				Type:        graphql.NewList(types.VehicleType),
				Description: "Get Vehicle List",
				Resolve:     resolver.VehicleListResolver,
			},
			"starship": &graphql.Field{
				Type:        types.StarshipType,
				Description: "Get Starship by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.StarshipResolver,
			},
			"starships": &graphql.Field{
				Type:        graphql.NewList(types.StarshipType),
				Description: "Get Starship List",
				Resolve:     resolver.StarshipListResolver,
			},
			"species": &graphql.Field{
				Type:        types.SpeciesType,
				Description: "Get Species by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.SpeciesResolver,
			},
			"speciesList": &graphql.Field{
				Type:        graphql.NewList(types.SpeciesType),
				Description: "Get Species List",
				Resolve:     resolver.SpeciesListResolver,
			},
			"planet": &graphql.Field{
				Type:        types.PlanetType,
				Description: "Get Planet by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.PlanetResolver,
			},
			"planets": &graphql.Field{
				Type:        graphql.NewList(types.PlanetType),
				Description: "Get Planet List",
				Resolve:     resolver.PlanetListResolver,
			},
			"character": &graphql.Field{
				Type:        types.CharacterType,
				Description: "Get Character by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.CharacterResolver,
			},
			"characters": &graphql.Field{
				Type:        graphql.NewList(types.CharacterType),
				Description: "Get Character List",
				Resolve:     resolver.CharacterListResolver,
			},
		},
	},
)
