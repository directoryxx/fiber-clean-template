package controller

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"

	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/encrypt"
	"github.com/directoryxx/fiber-clean-template/app/utils/jwt"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var initval *validator.Validate

// A UserController belong to the interface layer.
type UserController struct {
	Userservice service.UserService
	Logger      interfaces.Logger
}

func NewUserController(logger interfaces.Logger) *UserController {
	return &UserController{
		Userservice: service.UserService{
			UserRepository: repository.UserRepository{
				// SQLHandler:   sqlHandler,
				Ctx: context.Background(),
			},
		},
		Logger: logger,
	}
}

// Register
// @Summary Register user
// @Description Register user
// @Tags Authentication
// @Accept application/json
// @Param register body rules.RegisterValidation true "Register"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func (controller UserController) Register(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())
	var register *domain.User

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

	register.Password, _ = encrypt.CreateHash(register.Password, encrypt.DefaultParams)
	register.RoleID = 1

	data := controller.Userservice.CreateUser(register)

	token, errToken := jwt.CreateToken(uint(data.ID))

	if errToken != nil {
		controller.Logger.LogError("%s", errToken)
	}

	persistToken := jwt.CreateAuth(controller.Userservice, uint(data.ID), token)

	if persistToken != nil {
		controller.Logger.LogError("%s", errToken)
	}

	return c.JSON(&response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendaftar",
		Data: &response.RegisterResponse{
			Name:     data.Name,
			Username: data.Username,
			Token:    token.AccessToken,
		},
	})

}

// Login
// @Summary Login user
// @Description Login user
// @Tags Authentication
// @Accept application/json
// @Param login body rules.LoginValidation true "Login"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (controller UserController) Login(c *fiber.Ctx) error {

	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())
	var login *rules.LoginValidation
	err := c.BodyParser(&login)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	initval = validator.New()
	loginValidation(initval, controller.Userservice)
	errVal := initval.Struct(login)

	if errVal != nil {
		message := make(map[string]string)
		message["password"] = "Password harus lebih dari 6 karakter"
		errorResponse := validation.ValidateRequest(errVal, message)
		return c.JSON(errorResponse)
	}

	res := controller.Userservice.CheckUsername(login.Username)

	if res.ID == 0 {
		c.Status(422)
		err = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Data tidak ditemukan",
		})
		return err
	}

	check, _ := encrypt.ComparePasswordAndHash(login.Password, string(res.Password))

	if check {
		td, errToken := jwt.CreateToken(uint(res.ID))
		if errToken != nil {
			controller.Logger.LogError("%s", errToken)
		}

		jwt.CreateAuth(controller.Userservice, uint(res.ID), td)

		return c.JSON(&response.SuccessResponse{
			Message: "Berhasil login",
			Success: true,
			Data: &response.LoginResponse{
				Name:     res.Name,
				Username: res.Username,
				Token:    td.AccessToken,
			},
		})
	} else {
		c.Status(401)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Username/Password salah",
		})
	}

}

func registerValidation(initval *validator.Validate, service service.UserService) {
	initval.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return IsValidUsername(service, fl.Field().String())
	})
	initval.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return IsValidPassword(service, fl.Field().String())
	})
}

func IsValidPassword(service service.UserService, input string) bool {
	return len(input) > 6
}

func IsValidUsername(service service.UserService, input string) bool {
	count := service.CheckUsernameCount(input)

	return count == int64(0)
}

func loginValidation(initval *validator.Validate, service service.UserService) {
	initval.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return IsValidPassword(service, fl.Field().String())
	})
}
