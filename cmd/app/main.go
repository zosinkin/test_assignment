package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_pgx_pool "github.com/zosinkin/test_assignment.git/internal/core/repository/postgres/pool/pgx"
	core_http_server "github.com/zosinkin/test_assignment.git/internal/core/server"
	core_http_middleware "github.com/zosinkin/test_assignment.git/internal/core/transport/http/middleware"
	subscription_service "github.com/zosinkin/test_assignment.git/internal/features/subscriptions/service"
	subscription_transport_http "github.com/zosinkin/test_assignment.git/internal/features/subscriptions/transport/http"
	subscriptions_postgres_repository "github.com/zosinkin/test_assignment.git/internal/features/subscriptions/repository/postgres"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()


	logger.Debug("Initializing postgres connection pool")

	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()


	subRepo := subscriptions_postgres_repository.NewSubRepository(pool)
	subService := subscription_service.NewSubService(subRepo)

	subTransportHTTP := subscription_transport_http.NewSubHTTPHandler(subService)

	logger.Debug("Initializing HTTP server")
		httpServer := core_http_server.NewHTTPServer(
			core_http_server.NewConfigMust(),
			logger,
			core_http_middleware.RequestID(),
			core_http_middleware.Logger(logger),
			core_http_middleware.Trace(),
			core_http_middleware.Panic(),
		)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(subTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)


	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP Server run error", zap.Error(err))
	}



}