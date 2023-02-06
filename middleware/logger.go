package middleware

import (
	"context"
	"net/http"

	"github.com/iammunir/starwars-graphql-api/constant"
	"github.com/iammunir/starwars-graphql-api/logger"
)

// middleware for setting up Logger to be available accross request
func LoggerMiddleware(logger logger.Logger, nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		augmentedCtx := context.WithValue(ctx, constant.LoggerKey, logger)
		r = r.WithContext(augmentedCtx)
		nextHandler.ServeHTTP(w, r)
	})
}
