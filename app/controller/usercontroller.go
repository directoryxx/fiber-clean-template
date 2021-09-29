package controller

import (
	"context"
	"fmt"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type UserController struct {
	Userservice service.UserService
	Logger      interfaces.Logger
}

func NewUserController(sqlHandler *gen.Client, logger interfaces.Logger) *UserController {
	return &UserController{
		Userservice: service.UserService{
			UserRepository: repository.UserRepository{
				SQLHandler: sqlHandler,
				Ctx:        context.Background(),
			},
		},
		Logger: logger,
	}
}

func (controller UserController) IndexUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())
		var m map[string]string

		c.BodyParser(&m)

		controller.Userservice.CreateUser(m)

		fmt.Println(m)

		return c.JSON(fiber.Map{
			"message": "Hello World",
			// "data":    data,
		})
	}
}
