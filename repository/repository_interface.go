package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetVehicleList(ctx context.Context, selectQuery string) ([]entity.Vehicle, error)
	GetVehicleById(ctx context.Context, vehicleId int, selectQuery string) (*entity.Vehicle, error)

	GetStarshipList(ctx context.Context, selectQuery string) ([]entity.Starship, error)
	GetStarshipById(ctx context.Context, starshipId int, selectQuery string) (*entity.Starship, error)

	GetPlanetList(ctx context.Context, selectQuery string) ([]entity.Planet, error)
	GetPlanetById(ctx context.Context, planetId int, selectQuery string) (*entity.Planet, error)

	GetSpeciesList(ctx context.Context, selectQuery string) ([]entity.Species, error)
	GetSpeciesById(ctx context.Context, speciesId int, selectQuery string) (*entity.Species, error)
	GetSpeciesByPlanetIds(ctx context.Context, planetIds []int, selectQuery string) ([]entity.Species, error)

	GetCharacterList(ctx context.Context, selectQuery string) ([]entity.Character, error)
	GetCharacterById(ctx context.Context, characterId int, selectQuery string) (*entity.Character, error)
	GetCharacterListByPlanetIds(ctx context.Context, planetIds []int, selectQuery string) ([]entity.Character, error)
	GetCharacterListBySpeciesIds(ctx context.Context, speciesIds []int, selectQuery string) ([]entity.Character, error)
}

type repository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewRepository(db *gorm.DB, log *logrus.Logger) Repository {
	return &repository{
		db: db,
	}
}
