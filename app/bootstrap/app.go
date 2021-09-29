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

	// repository.NewRepository(sqlHandler, ctx)

	routes.RegisterRoute(app, sqlHandler, ctx, log)

	errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT"))

	if errApp != nil {
		log.LogError("%s", errApp)
	}

	// log.Fatal()
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
