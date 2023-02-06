package resolver

import (
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/repository"
)

var SpeciesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: SpeciesResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	id, ok := p.Args["id"].(int)
	if !ok {
		log.Error("error casting id value")
		return nil, constant.ErrCastingValue
	}

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Species{}, selectedFields)

	result, err := repo.GetSpeciesById(p.Context, id, selectQuery)
	if err != nil {
		log.Error("error getting vehicle by id")
	}

	return result, err
}

var SpeciesListResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: SpeciesListResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Species{}, selectedFields)

	result, err := repo.GetSpeciesList(p.Context, selectQuery)
	if err != nil {
		log.Error("error getting vehicle by id")
	}

	return result, err
}

var SpeciesPlanetResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: SpeciesPlanetResolver")

	species, ok := p.Source.(*entity.Species)
	if !ok {
		return nil, fmt.Errorf("error getting species parent value")
	}

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Species{}, selectedFields)

	loaders := p.Context.Value(constant.LoaderKey).(map[string]*dataloader.Loader)
	planetLoader := loaders[constant.PlanetLoaderKey]

	key := NewResolverKey(fmt.Sprintf("%d", species.HomeworldId), selectQuery)

	ch := make(chan resultLoader, 1)
	go func() {
		loaderResult, err := planetLoader.Load(p.Context, key)()
		if err != nil {
			ch <- resultLoader{err: err}
		}
		planet := loaderResult.(*entity.Planet)
		ch <- resultLoader{
			data: planet,
			err:  nil,
		}
		close(ch)
	}()

	return func() (interface{}, error) {
		r := <-ch
		if r.err != nil {
			return nil, r.err
		}
		return r.data, nil
	}, nil
}
