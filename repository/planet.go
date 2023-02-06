package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"gorm.io/gorm"
)

func (r *repository) GetPlanetList(ctx context.Context, selectQuery string) ([]*entity.Planet, error) {
	r.log.Trace("Enter: repository GetPlanetList")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var planetList []*entity.Planet
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Planet{}).
		Scan(&planetList).
		Error

	if errDb != nil {
		r.log.Error("database error when getting planet list")
		return nil, errDb
	}

	if len(planetList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return planetList, nil
}

func (r *repository) GetPlanetById(ctx context.Context, planetId int, selectQuery string) (*entity.Planet, error) {

	r.log.Trace("Enter: repository GetPlanetById")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var planet entity.Planet
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Planet{}).
		Where("id = ?", planetId).
		Scan(&planet).
		Error

	if errDb != nil {
		r.log.Error("database error when getting planet data")
		return nil, errDb
	}

	return &planet, nil
}

func (r *repository) GetPlanetListByIds(ctx context.Context, planetIds []int, selectQuery string) ([]*entity.Planet, error) {

	r.log.Trace("Enter: repository GetPlanetListByIds")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var planetList []*entity.Planet
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Planet{}).
		Where("id IN (?)", planetIds).
		Scan(&planetList).
		Error

	if errDb != nil {
		r.log.Error("database error when getting planet data")
		return nil, errDb
	}

	if len(planetList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return planetList, nil
}
