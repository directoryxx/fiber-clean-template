package main

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/bootstrap"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
)

func main() {
	ctx := context.Background()
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	_, enforcer, err := infrastructure.NewSQLHandler(ctx)
	if err != nil {
		logger.LogError("%s", err)
	}

	redisHandler := infrastructure.RedisInit()

	bootstrap.Dispatch(ctx, logger, redisHandler, enforcer)
}
