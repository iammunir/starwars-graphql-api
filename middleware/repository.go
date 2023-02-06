package middleware

import (
	"context"
	"net/http"

	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/repository"
)

// middleware for setting up Repository to be available accross request
func RepositoryMiddleware(repo repository.Repository, nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		augmentedCtx := context.WithValue(ctx, constant.RepositoryKey, repo)
		r = r.WithContext(augmentedCtx)
		nextHandler.ServeHTTP(w, r)
	})
}
