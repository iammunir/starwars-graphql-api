package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/graphql/resolver"
	"github.com/iammunir/starwars-graphql-api/graphql/schema"
	"github.com/iammunir/starwars-graphql-api/graphql/types"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/middleware"
	"github.com/iammunir/starwars-graphql-api/repository"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	types.SpeciesType.AddFieldConfig("homeworld", &graphql.Field{
		Type:    types.PlanetType,
		Resolve: resolver.SpeciesPlanetResolver,
	})
}

func InitApp(dbConn *gorm.DB, log logger.Logger) *http.ServeMux {

	timeoutThreshold, err := time.ParseDuration(viper.GetString("TIMEOUT"))
	if err != nil {
		log.Fatal("error setting timeout threshold: %s", err.Error())
	}

	definedSchema, err := schema.NewGraphQLSchema()
	if err != nil {
		log.Fatal("error defining schema: %s", err.Error())
	}

	h := handler.New(&handler.Config{
		Schema:   &definedSchema,
		Pretty:   true,
		GraphiQL: true,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			return map[string]interface{}{
				constant.REQUEST_START_TIME: time.Now(),
			}
		},
		FormatErrorFn: func(err error) gqlerrors.FormattedError {
			if err == context.DeadlineExceeded {
				err = errors.New("processing time is too long")
			}

			return gqlerrors.FormattedError{
				Message: err.Error(),
			}
		},
		ResultCallbackFn: func(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {

			reqStart := params.RootObject[constant.REQUEST_START_TIME].(time.Time)

			var queryName string
			if queryNameObj := params.RootObject[constant.QUERY_NAME]; queryNameObj != nil {
				queryName = queryNameObj.(string)
			}

			var err string
			if result.HasErrors() {
				err = "with error"
			} else {
				err = "without error"
			}
			log.Info(fmt.Sprintf("request %s has been responded %s for %s", queryName, err, time.Since(reqStart)))
		},
	})
	log.Debug("graphql handler was succefully initialized")

	mux := http.NewServeMux()

	repo := repository.NewRepository(dbConn, log)

	graphqlHandler := middleware.LoggerMiddleware(log, h)
	graphqlHandler = middleware.RepositoryMiddleware(repo, graphqlHandler)
	graphqlHandler = middleware.TimeoutMiddleware(timeoutThreshold, graphqlHandler) // second middleware
	graphqlHandler = middleware.CORSMiddleware(graphqlHandler)                      // first middleware
	dataloderMiddleware := middleware.Dataloader(repo)

	mux.Handle("/graphql", dataloderMiddleware(graphqlHandler))
	mux.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong")
	}))

	return mux

}
