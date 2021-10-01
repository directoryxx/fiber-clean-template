package bootstrap

import (
	"context"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/routes"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// Dispatch is handle routing
func Dispatch(sqlHandler *gen.Client, ctx context.Context, log interfaces.Logger, redisHandler *redis.Client, enforcer *casbin.Enforcer) {
	app := fiber.New()

	routes.RegisterRoute(app, sqlHandler, ctx, log, redisHandler, enforcer)

	errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT"))

	if errApp != nil {
		log.LogError("%s", errApp)
	}
}
