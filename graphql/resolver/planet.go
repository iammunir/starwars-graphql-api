package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/repository"
)

var PlanetResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: PlanetResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	id, ok := p.Args["id"].(int)
	if !ok {
		log.Error("error casting cifno value")
		return nil, constant.ErrCastingValue
	}

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Planet{}, selectedFields)

	result, err := repo.GetPlanetById(p.Context, id, selectQuery)
	if err != nil {
		log.Error("error getting vehicle by id")
	}

	return result, err
}

var PlanetListResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: PlanetListResolver")

	rootValue := p.Info.RootValue.(map[string]interface{})
	fieldName := p.Info.FieldName
	rootValue[constant.QUERY_NAME] = fieldName

	repo := p.Context.Value(constant.RepositoryKey).(repository.Repository)

	selectedFields, err := getSelectedFields(p)
	if err != nil {
		log.Debug(fmt.Sprintf("error getting selected fields: %s", err.Error()))
	}
	selectQuery := GetColumnListFromAttributes(entity.Planet{}, selectedFields)

	result, err := repo.GetPlanetList(p.Context, selectQuery)
	if err != nil {
		log.Error("error getting vehicle by id")
	}

	return result, err
}
