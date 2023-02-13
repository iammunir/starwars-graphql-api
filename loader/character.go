package loader

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/graphql/resolver"
	"github.com/iammunir/starwars-graphql-api/repository"
)

func NewCharacterDataloader(repo repository.Repository) *dataloader.Loader {
	batchFunc := func(c context.Context, keys dataloader.Keys) []*dataloader.Result {

		handleError := func(err error) []*dataloader.Result {
			var results []*dataloader.Result
			var result dataloader.Result
			result.Error = err
			results = append(results, &result)
			return results
		}

		var results []*dataloader.Result
		var characterIds []int
		for _, key := range keys {
			id, err := strconv.ParseInt(key.String(), 10, 32)
			if err != nil {
				return handleError(err)
			}
			characterIds = append(characterIds, int(id))
		}

		selectQuery := keys[0].(*resolver.ResolverKey).GetSelectQuery()

		characters, err := repo.GetCharacterListByIds(c, characterIds, selectQuery)
		if err != nil {
			return handleError(err)
		}

		var charactersMap = make(map[int]*entity.Character)
		charactersMap[0] = nil // assign nil as default value for character key NULL

		for _, character := range characters {
			if _, found := charactersMap[character.Id]; !found {
				charactersMap[character.Id] = character
			}
		}

		for _, characterId := range characterIds {
			character, ok := charactersMap[characterId]
			if !ok {
				return nil
			}
			r := dataloader.Result{
				Data:  character,
				Error: nil,
			}
			results = append(results, &r)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFunc)

	return loader
}

func NewCharacterBySpeciesDataloader(repo repository.Repository) *dataloader.Loader {
	batchFunc := func(c context.Context, keys dataloader.Keys) []*dataloader.Result {

		handleError := func(err error) []*dataloader.Result {
			var results []*dataloader.Result
			var result dataloader.Result
			result.Error = err
			results = append(results, &result)
			return results
		}

		var results []*dataloader.Result
		var speciesIds []int
		for _, key := range keys {
			id, err := strconv.ParseInt(key.String(), 10, 32)
			if err != nil {
				return handleError(err)
			}
			speciesIds = append(speciesIds, int(id))
		}

		selectQuery := keys[0].(*resolver.ResolverKey).GetSelectQuery()

		characters, err := repo.GetCharacterListBySpeciesIds(c, speciesIds, selectQuery)
		if err != nil {
			return handleError(err)
		}

		var charactersMap = make(map[int][]*entity.Character)
		charactersMap[0] = nil // assign nil as default value for character key NULL

		for _, character := range characters {
			if _, found := charactersMap[character.SpeciesId]; !found {
				charactersMap[character.SpeciesId] = []*entity.Character{
					character,
				}
			} else {
				charactersMap[character.SpeciesId] = append(charactersMap[character.SpeciesId], character)
			}
		}

		for _, speciesId := range speciesIds {
			character, ok := charactersMap[speciesId]
			if !ok {
				return nil
			}
			r := dataloader.Result{
				Data:  character,
				Error: nil,
			}
			results = append(results, &r)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFunc)

	return loader
}
