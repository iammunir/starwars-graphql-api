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

var CharacterResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: CharacterResolver")

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
	selectQuery := GetColumnListFromAttributes(entity.Character{}, selectedFields)

	result, err := repo.GetCharacterById(p.Context, id, selectQuery)
	if err != nil {
		log.Error("error getting character by id")
	}

	return result, err
}

var CharacterListResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: CharacterListResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Character{}, selectedFields)

	result, err := repo.GetCharacterList(p.Context, selectQuery)
	if err != nil {
		log.Error("error getting character by id")
	}

	return result, err
}

var CharacterPlanetResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: CharacterPlanetResolver")

	character, ok := p.Source.(*entity.Character)
	if !ok {
		return nil, fmt.Errorf("error getting character parent value")
	}

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectedFields = append(selectedFields, "id") // id must be included in selected fields to avoid error
	selectQuery := GetColumnListFromAttributes(entity.Planet{}, selectedFields)

	loaders := p.Context.Value(constant.LoaderKey).(map[string]*dataloader.Loader)
	planetLoader := loaders[constant.PlanetLoaderKey]

	key := NewResolverKey(fmt.Sprintf("%d", character.HomeworldId), selectQuery)

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

var CharacterSpeciesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: CharacterSpeciesResolver")

	character, ok := p.Source.(*entity.Character)
	if !ok {
		return nil, fmt.Errorf("error getting character parent value")
	}

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectedFields = append(selectedFields, "id") // id must be included in selected fields to avoid error
	selectQuery := GetColumnListFromAttributes(entity.Species{}, selectedFields)

	loaders := p.Context.Value(constant.LoaderKey).(map[string]*dataloader.Loader)
	speciesLoader := loaders[constant.SpeciesLoaderKey]

	key := NewResolverKey(fmt.Sprintf("%d", character.SpeciesId), selectQuery)

	ch := make(chan resultLoader, 1)
	go func() {
		loaderResult, err := speciesLoader.Load(p.Context, key)()
		if err != nil {
			ch <- resultLoader{err: err}
		}
		species := loaderResult.(*entity.Species)
		ch <- resultLoader{
			data: species,
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
