package controller

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"

	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/validation"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var initval *validator.Validate

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

func (controller UserController) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())
		var register *rules.RegisterValidation

		errRequest := c.BodyParser(&register)

		if errRequest != nil {
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Invalid request body",
			})
		}

		initval = validator.New()
		registerValidation(initval, controller.Userservice)
		errVal := initval.Struct(register)

		if errVal != nil {
			message := make(map[string]string)
			message["username"] = "Username telah terdaftar"
			message["password"] = "Password harus lebih dari 6 karakter"
			errorResponse := validation.ValidateRequest(errVal, message)
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errorResponse,
			})
		}

		data, err := controller.Userservice.CreateUser(register)

		if err != nil {
			controller.Logger.LogError("%s", err)
		}

		return c.JSON(fiber.Map{
			"message": "Hello World",
			"data":    data,
		})
	}
}

func registerValidation(initval *validator.Validate, service service.UserService) {
	// initval.RegisterValidation("username", func(fl validator.FieldLevel) bool {
	// 	return IsValidUsername(service, fl.Field().String())
	// })
	initval.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return IsValidPassword(service, fl.Field().String())
	})
}

func IsValidPassword(service service.UserService, input string) bool {
	return len(input) > 6
}

// func IsValidUsername(service service.UserService, input string) bool {
// 	// count := service.GetUsernameCount(input)
// 	// return count == int64(0)
// }
