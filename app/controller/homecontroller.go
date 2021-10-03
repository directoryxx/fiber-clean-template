package controller

import (
	"context"
	"fmt"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/jwt"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/session"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type HomeController struct {
	Userservice service.UserService
	Logger      interfaces.Logger
}

func NewHomeController(sqlHandler *gen.Client, logger interfaces.Logger, redisHandler *redis.Client) *HomeController {
	return &HomeController{
		Userservice: service.UserService{
			UserRepository: repository.UserRepository{
				SQLHandler:   sqlHandler,
				Ctx:          context.Background(),
				RedisHandler: redisHandler,
			},
		},
		Logger: logger,
	}
}

func (controller HomeController) Current() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := session.GetAuth()
		fmt.Println(auth)

		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		token, err := jwt.ExtractTokenMetadata(c)
		if err != nil {
			controller.Logger.LogError("%s", err)
		}

		res, errGet := controller.Userservice.CurrentUser(token.UserId)

		if errGet != nil {
			controller.Logger.LogError("%s", errGet)
		}

		return c.JSON(&response.CurrentResponse{
			Name:     res.Name,
			Username: res.Username,
		})
	}
}
