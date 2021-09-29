package bootstrap

import (
	"context"
	"os"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/routes"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/gofiber/fiber/v2"
)

// Dispatch is handle routing
func Dispatch(sqlHandler *gen.Client, ctx context.Context, log interfaces.Logger) {
	app := fiber.New()

	routes.RegisterRoute(app, sqlHandler, ctx, log)

	errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT"))

	if errApp != nil {
		log.LogError("%s", errApp)
	}
}
