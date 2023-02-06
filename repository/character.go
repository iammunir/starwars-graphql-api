package repository

import (
	"context"

	"github.com/iammunir/starwars-graphql-api/entity"
	"gorm.io/gorm"
)

func (r *repository) GetCharacterList(ctx context.Context, selectQuery string) ([]*entity.Character, error) {
	r.log.Trace("Enter: repository GetCharacterList")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var characterList []*entity.Character
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Character{}).
		Scan(&characterList).
		Error

	if errDb != nil {
		r.log.Error("database error when getting character list")
		return nil, errDb
	}

	if len(characterList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return characterList, nil
}

func (r *repository) GetCharacterById(ctx context.Context, characterId int, selectQuery string) (*entity.Character, error) {

	r.log.Trace("Enter: repository GetCharacterById")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var character entity.Character
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Character{}).
		Where("id = ?", characterId).
		Scan(&character).
		Error

	if errDb != nil {
		r.log.Error("database error when getting character data")
		return nil, errDb
	}

	return &character, nil
}

func (r *repository) GetCharacterListByPlanetIds(ctx context.Context, planetIds []int, selectQuery string) ([]*entity.Character, error) {
	r.log.Trace("Enter: repository GetCharacterListByPlanetIds")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var characterList []*entity.Character
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Character{}).
		Where("homeworld_id IN (?)", planetIds).
		Scan(&characterList).
		Error

	if errDb != nil {
		r.log.Error("database error when getting character list by planet ids")
		return nil, errDb
	}

	if len(characterList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return characterList, nil
}

func (r *repository) GetCharacterListBySpeciesIds(ctx context.Context, speciesIds []int, selectQuery string) ([]*entity.Character, error) {
	r.log.Trace("Enter: repository GetCharacterListBySpeciesIds")

	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var characterList []*entity.Character
	errDb := r.db.
		Select(selectQuery).
		Model(&entity.Character{}).
		Where("species_id IN (?)", speciesIds).
		Scan(&characterList).
		Error

	if errDb != nil {
		r.log.Error("database error when getting character list by species ids")
		return nil, errDb
	}

	if len(characterList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return characterList, nil
}
