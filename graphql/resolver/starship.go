package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/repository"
)

var StarshipResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: StarshipResolver")

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
	selectQuery := GetColumnListFromAttributes(entity.Starship{}, selectedFields)

	result, err := repo.GetStarshipById(p.Context, id, selectQuery)
	if err != nil {
		log.Error("error getting starship by id")
	}

	return result, err
}

var StarshipListResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: StarshipListResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Starship{}, selectedFields)

	result, err := repo.GetStarshipList(p.Context, selectQuery)
	if err != nil {
		log.Error("error getting starship by id")
	}

	return result, err
}
