package main

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/bootstrap"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
)

func main() {
	// if ()
	ctx := context.Background()
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	enforcer, err := infrastructure.NewSQLHandler(ctx)
	if err != nil {
		logger.LogError("%s", err)
	}

	// redisHandler := infrastructure.RedisInit()

	bootstrap.Dispatch(ctx, logger, enforcer)
}
