package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	core_http_middleware "github.com/glebateee/taskapp/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	cfg    Config
	logger *core_logger.Logger

	middlewares []core_http_middleware.Middleware
}

func NewHTTPServer(
	cfg Config,
	logger *core_logger.Logger,
	middlewares ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:         http.NewServeMux(),
		cfg:         cfg,
		logger:      logger,
		middlewares: middlewares,
	}
}
func (s *HTTPServer) RegisterApiRouters(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := fmt.Sprintf("/api/%s", router.apiVersion)
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleware()))
	}
}

func (s *HTTPServer) Run(
	ctx context.Context,
) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middlewares...)
	server := &http.Server{
		Addr:    s.cfg.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.logger.Debug("start HTTP server", zap.String("addr", s.cfg.Addr))
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		s.logger.Warn("shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.cfg.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
		s.logger.Warn("HTTP server stopped")
	}
	return nil
}
