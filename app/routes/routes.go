package routes

import (
	"context"

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

	app.Get("/dashboard", monitor.New())

	app.Post("/register", UserController.Register())
	app.Post("/login", UserController.Login())

	app.Use(middleware.JWTProtected(service.UserService{
		UserRepository: repository.UserRepository{
			// SQLHandler:   sqlHandler,
			Ctx: ctx,
		},
	}))

	enforcer.AddPolicy("admin", "role", "manage")
	enforcer.AddPolicy("admin", "permission", "manage")

	HomeController.HomeRouter()
	RoleController.RoleRouter()
	PermissionController.PermissionRouter()

}
