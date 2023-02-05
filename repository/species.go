package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"gorm.io/gorm"
)

func (r *repository) GetSpeciesList(ctx context.Context, selectQuery string) ([]entity.Species, error) {
	r.log.Trace("Enter: repository GetSpeciesList")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var speciesList []entity.Species
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Species{}).
		Scan(&speciesList).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting species list")
		return nil, errDb
	}

	if len(speciesList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return speciesList, nil
}

func (r *repository) GetSpeciesById(ctx context.Context, speciesId int, selectQuery string) (*entity.Species, error) {

	r.log.Trace("Enter: repository GetSpeciesById")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var species entity.Species
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Species{}).
		Where("id = ?", speciesId).
		Scan(&species).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting species data")
		return nil, errDb
	}

	return &species, nil
}

func (r *repository) GetSpeciesByPlanetIds(ctx context.Context, planetIds []int, selectQuery string) ([]entity.Species, error) {
	r.log.Trace("Enter: repository GetSpeciesByPlanetIds")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var speciesList []entity.Species
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Species{}).
		Where("homeworld_id IN (?)", planetIds).
		Scan(&speciesList).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting species list by planet ids")
		return nil, errDb
	}

	if len(speciesList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return speciesList, nil
}
