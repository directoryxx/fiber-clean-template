package routes

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/controller"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, sqlHandler *gen.Client, ctx context.Context, log interfaces.Logger, redisHandler *redis.Client) {
	UserController := controller.NewUserController(sqlHandler, log, redisHandler)

	app.Post("/register", UserController.Register())
	app.Post("/login", UserController.Login())
}
