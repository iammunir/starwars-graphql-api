package loader

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/graphql/resolver"
	"github.com/iammunir/starwars-graphql-api/repository"
)

func NewSpeciesDataloader(repo repository.Repository) *dataloader.Loader {
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

		species, err := repo.GetSpeciesListByIds(c, speciesIds, selectQuery)
		if err != nil {
			return handleError(err)
		}

		var speciesMap = make(map[int]*entity.Species)
		speciesMap[0] = nil // assign nil as default value for species key NULL

		for _, species := range species {
			if _, found := speciesMap[species.Id]; !found {
				speciesMap[species.Id] = species
			}
		}

		for _, speciesId := range speciesIds {
			species, ok := speciesMap[speciesId]
			if !ok {
				return nil
			}
			r := dataloader.Result{
				Data:  species,
				Error: nil,
			}
			results = append(results, &r)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFunc)

	return loader
}
