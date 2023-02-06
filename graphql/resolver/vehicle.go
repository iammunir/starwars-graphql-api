package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/entity"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/repository"
)

var VehicleResolver = func(p graphql.ResolveParams) (interface{}, error) {
	log := p.Context.Value(constant.LoggerKey).(logger.Logger)
	log.Trace("Enter: VehicleResolver")

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
	selectQuery := GetColumnListFromAttributes(entity.Vehicle{}, selectedFields)

	result, err := repo.GetVehicleById(p.Context, id, selectQuery)
	if err != nil {
		log.Error("error getting vehicle by id")
	}

	return result, err
}
