package main

import (
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// log.Fatal(app.Listen(":"+os.Getenv("APP_PORT")))
}
