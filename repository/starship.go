package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"gorm.io/gorm"
)

func (r *repository) GetStarshipList(ctx context.Context, selectQuery string) ([]entity.Starship, error) {
	r.log.Trace("Enter: repository GetStarshipList")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var starshipList []entity.Starship
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Vehicle{}).
		Scan(&starshipList).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting starship list")
		return nil, errDb
	}

	if len(starshipList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return starshipList, nil
}

func (r *repository) GetStarshipById(ctx context.Context, starshipId int, selectQuery string) (*entity.Starship, error) {

	r.log.Trace("Enter: repository GetStarshipById")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var starship entity.Starship
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Starship{}).
		Where("id = ?", starshipId).
		Scan(&starship).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting starship data")
		return nil, errDb
	}

	return &starship, nil
}
