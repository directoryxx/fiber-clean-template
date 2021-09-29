package main

import (
	"context"
	"log"

	"github.com/directoryxx/fiber-clean-template/app/bootstrap"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
)

func main() {
	ctx := context.Background()
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bootstrap.Dispatch(sqlHandler, ctx, logger)
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// log.Fatal(app.Listen(":"+os.Getenv("APP_PORT")))
}
