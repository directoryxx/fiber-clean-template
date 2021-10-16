package routes

import (
	"context"
	"io/ioutil"
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/controller"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/middleware"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func RegisterRoute(app *fiber.App, ctx context.Context, log interfaces.Logger, enforcer *casbin.Enforcer) {
	UserController := controller.NewUserController(log)
	HomeController := controller.NewHomeController(log, app)
	RoleController := controller.NewRoleController(log, app, enforcer)
	PermissionController := controller.NewPermissionController(log, enforcer, app)
	QueueController := controller.NewQueueController(log)
	UploadController := controller.NewUploadController(log, app, enforcer)

	app.Get("/docs/*", swagger.New(
		swagger.Config{
			URL:         os.Getenv("APP_URL") + "/swagger.json",
			DeepLinking: false,
			// Expand ("list") or Collapse ("none") tag groups by default
			DocExpansion: "none",
		},
	)) // default

	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		body, err := ioutil.ReadFile("docs/swagger.json")
		if err != nil {
			panic(err)
		}
		return c.SendString(string(body))
	})

	app.Get("/dashboard", monitor.New())

	app.Post("/register", UserController.Register)
	app.Post("/login", UserController.Login)

	app.Use(middleware.JWTProtected(service.UserService{
		UserRepository: repository.UserRepository{
			Ctx: ctx,
		},
	}))

	enforcer.AddPolicy("admin", "role", "manage")
	enforcer.AddPolicy("admin", "permission", "manage")
	enforcer.AddPolicy("admin", "upload", "manage")

	HomeController.HomeRouter()
	RoleController.RoleRouter()
	PermissionController.PermissionRouter()
	UploadController.UploadRouter()

	app.Get("/test", QueueController.TestQueue)
}
