package loader

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/graphql/resolver"
	"github.com/iammunir/starwars-graphql-api/repository"
)

func NewPlanetDataloader(repo repository.Repository) *dataloader.Loader {
	batchFunc := func(c context.Context, keys dataloader.Keys) []*dataloader.Result {

		handleError := func(err error) []*dataloader.Result {
			var results []*dataloader.Result
			var result dataloader.Result
			result.Error = err
			results = append(results, &result)
			return results
		}

		var results []*dataloader.Result
		var planetIds []int
		for _, key := range keys {
			id, err := strconv.ParseInt(key.String(), 10, 32)
			if err != nil {
				return handleError(err)
			}
			planetIds = append(planetIds, int(id))
		}

		selectQuery := keys[0].(*resolver.ResolverKey).GetSelectQuery()

		planets, err := repo.GetPlanetListByIds(c, planetIds, selectQuery)
		if err != nil {
			return handleError(err)
		}

		var planetsMap = make(map[int]*entity.Planet)
		planetsMap[0] = nil // assign nil as default value for planet key NULL

		for _, planet := range planets {
			if _, found := planetsMap[planet.Id]; !found {
				planetsMap[planet.Id] = planet
			}
		}

		for _, planetId := range planetIds {
			planet, ok := planetsMap[planetId]
			if !ok {
				return nil
			}
			r := dataloader.Result{
				Data:  planet,
				Error: nil,
			}
			results = append(results, &r)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFunc)

	return loader
}
