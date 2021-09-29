package routes

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/controller"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, sqlHandler *gen.Client, ctx context.Context, log interfaces.Logger) {
	UserController := controller.NewUserController(sqlHandler, log)
	
	app.Post("/register", UserController.Register())
}
