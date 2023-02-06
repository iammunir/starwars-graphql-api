package middleware

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeMilis time.Duration, nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), timeMilis)
		defer cancel()
		req = req.WithContext(ctx)
		nextHandler.ServeHTTP(w, req)
	})
}
