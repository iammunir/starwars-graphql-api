package middleware

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/loader"
	"github.com/iammunir/starwars-graphql-api/repository"
)

// middleware for setting up Dataloader to be available accross request
func Dataloader(repo repository.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			planetDataloader := loader.NewPlanetDataloader(repo)

			var loaders = make(map[string]*dataloader.Loader, 1) // make sure size matches with quantity of loaders
			loaders[constant.PlanetLoaderKey] = planetDataloader

			augmentedCtx := context.WithValue(ctx, constant.LoaderKey, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}
