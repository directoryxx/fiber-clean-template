package controller

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/session"
	"github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type HomeController struct {
	Userservice service.UserService
	Logger      interfaces.Logger
	Fiber       *fiber.App
}

func NewHomeController(logger interfaces.Logger, fiber *fiber.App) *HomeController {
	return &HomeController{
		Userservice: service.UserService{
			UserRepository: repository.UserRepository{
				// SQLHandler:   sqlHandler,
				Ctx: context.Background(),
			},
		},
		Logger: logger,
		Fiber:  fiber,
	}
}

func (controller HomeController) HomeRouter() {
	controller.Fiber.Get("/current", controller.current)
}

// Current User
// @Summary Current User
// @Description Current User
// @Tags Authentication
// @Param Authorization header string true "With the bearer started"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /current [get]
func (controller HomeController) current(c *fiber.Ctx) error {
	auth := session.GetAuth()
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	// token, err := jwt.ExtractTokenMetadata(c)
	// if err != nil {
	// 	controller.Logger.LogError("%s", err)
	// }

	// res, errGet := controller.Userservice.CurrentUser(token.UserId)

	// if errGet != nil {
	// 	controller.Logger.LogError("%s", errGet)
	// }

	return c.JSON(&response.CurrentResponse{
		Name:     auth.Auth.Name,
		Username: auth.Auth.Username,
	})

}
