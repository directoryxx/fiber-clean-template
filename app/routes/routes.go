package routes

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/controller"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/middleware"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func RegisterRoute(app *fiber.App, ctx context.Context, log interfaces.Logger, redisHandler *redis.Client, enforcer *casbin.Enforcer) {
	UserController := controller.NewUserController(log, redisHandler)
	HomeController := controller.NewHomeController(log, redisHandler, app)
	RoleController := controller.NewRoleController(log, app)
	PermissionController := controller.NewPermissionController(log, enforcer, app)

	app.Get("/dashboard", monitor.New())

	app.Post("/register", UserController.Register())
	app.Post("/login", UserController.Login())

	app.Use(middleware.JWTProtected(service.UserService{
		UserRepository: repository.UserRepository{
			// SQLHandler:   sqlHandler,
			Ctx:          ctx,
			RedisHandler: redisHandler,
		},
	}))

	HomeController.HomeRouter()
	RoleController.RoleRouter()
	PermissionController.PermissionRouter()

}
