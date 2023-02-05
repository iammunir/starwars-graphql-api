package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"gorm.io/gorm"
)

func (r *repository) GetVehicleList(ctx context.Context, selectQuery string) ([]entity.Vehicle, error) {
	r.log.Trace("Enter: repository GetVehicleList")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var vehicleList []entity.Vehicle
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Vehicle{}).
		Scan(&vehicleList).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting vehicle list")
		return nil, errDb
	}

	if len(vehicleList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return vehicleList, nil
}

func (r *repository) GetVehicleById(ctx context.Context, vehicleId int, selectQuery string) (*entity.Vehicle, error) {

	r.log.Trace("Enter: repository GetVehicleById")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var vehicle entity.Vehicle
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Vehicle{}).
		Where("id = ?", vehicleId).
		Scan(&vehicle).
		Error

	if errDb != nil {
		r.log.WithError(errDb).Error("database error when getting vehicle data")
		return nil, errDb
	}

	return &vehicle, nil
}
