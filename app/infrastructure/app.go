package infrastructure

import (
	"log"
	"os"

	"github.com/directoryxx/fiber-clean-template/app/usecases"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/gofiber/fiber/v2"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, sqlHandler *gen.Client) {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
	// userController := interfaces.NewUserController(sqlHandler, logger)
	// postController := interfaces.NewPostController(sqlHandler, logger)

	// r := chi.NewRouter()
	// r.Get("/users", userController.Index)
	// r.Get("/user", userController.Show)
	// r.Get("/posts", postController.Index)
	// r.Post("/post", postController.Store)
	// r.Delete("/post", postController.Destroy)

	// if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
	// 	logger.LogError("%s", err)
	// }
}
