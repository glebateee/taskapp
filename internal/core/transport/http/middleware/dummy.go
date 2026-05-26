package core_http_middleware

import (
	"net/http"

	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	"go.uber.org/zap"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			logger := core_logger.FromContextMust(ctx)

			logger.Debug(
				"----> before",
				zap.String("str", s),
			)
			next.ServeHTTP(w, r)

			logger.Debug(
				"<---- after",
				zap.String("str", s),
			)
		})
	}
}
