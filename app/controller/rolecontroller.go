package controller

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type RoleController struct {
	Roleservice service.RoleService
	Logger      interfaces.Logger
}

func NewRoleController(logger interfaces.Logger) *RoleController {
	return &RoleController{
		Roleservice: service.RoleService{
			RoleRepository: repository.RoleRepository{
				// SQLHandler: sqlHandler,
				Ctx: context.Background(),
			},
		},
		Logger: logger,
	}
}

func (controller RoleController) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		// token, err := jwt.ExtractTokenMetadata(c)
		// if err != nil {
		// 	controller.Logger.LogError("%s", err)
		// }

		// // res, errGet := controller.Userservice.CurrentUser(token.UserId)

		// if errGet != nil {
		// 	controller.Logger.LogError("%s", errGet)
		// }

		return c.JSON(&response.CurrentResponse{
			Name:     "test",
			Username: "test",
		})
	}
}

func (controller RoleController) CreateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		var role *rules.RoleValidation
		errRequest := c.BodyParser(&role)

		if errRequest != nil {
			controller.Logger.LogError("%s", errRequest)
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errRequest,
			})
		}

		initval = validator.New()
		roleValidation(initval, controller.Roleservice)
		errVal := initval.Struct(role)

		if errVal != nil {
			message := make(map[string]string)
			message["name"] = "Role telah terdaftar"
			errorResponse := validation.ValidateRequest(errVal, message)
			return c.JSON(errorResponse)
		}

		dataRole, errCreate := controller.Roleservice.CreateRole(role)

		if errCreate != nil {
			controller.Logger.LogError("%s", errCreate)
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errCreate,
			})
		}

		return c.JSON(response.SuccessResponse{
			Success: true,
			Data:    dataRole,
			Message: "Role berhasil dibuat",
		})

	}
}

func (controller RoleController) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		var role *rules.RoleValidation
		errRequest := c.BodyParser(&role)

		id, err := c.ParamsInt("id")

		data := controller.Roleservice.GetById(id)

		if data == nil {
			c.Status(404)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Data tidak ditemukan",
			})
		}

		if err != nil {
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Silahkan periksa kembali",
			})
		}

		if errRequest != nil {
			controller.Logger.LogError("%s", errRequest)
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errRequest,
			})
		}

		initval = validator.New()
		roleValidation(initval, controller.Roleservice)
		errVal := initval.Struct(role)

		if errVal != nil {
			message := make(map[string]string)
			message["name"] = "Role telah terdaftar"
			errorResponse := validation.ValidateRequest(errVal, message)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errorResponse,
			})
		}

		dataRole, errCreate := controller.Roleservice.UpdateRole(id, role)

		if errCreate != nil {
			controller.Logger.LogError("%s", errCreate)
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: errCreate,
			})
		}

		return c.JSON(response.SuccessResponse{
			Success: true,
			Data:    dataRole,
			Message: "Role berhasil diubah",
		})

	}
}

func (controller RoleController) GetRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		id, err := c.ParamsInt("id")

		if err != nil {
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Silahkan periksa kembali",
			})
		}

		roleData := controller.Roleservice.GetById(id)

		if roleData.ID == 0 {
			c.Status(404)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Data tidak ditemukan",
			})
		}

		return c.JSON(response.SuccessResponse{
			Success: true,
			Data:    roleData,
			Message: "Role berhasil diambil",
		})

	}
}

func (controller RoleController) DeleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

		id, err := c.ParamsInt("id")

		if err != nil {
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Silahkan periksa kembali",
			})
		}

		deleteRole := controller.Roleservice.DeleteRole(id)

		if deleteRole != nil {
			c.Status(422)
			return c.JSON(&response.ErrorResponse{
				Success: false,
				Message: "Data gagal dihapus",
			})
		}

		return c.JSON(response.SuccessResponse{
			Success: true,
			// Data:    roleData,
			Message: "Role berhasil dihappus",
		})

	}
}

func roleValidation(initval *validator.Validate, service service.RoleService) {
	initval.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		return uniqueRole(service, fl.Field().String())
	})
}

func uniqueRole(service service.RoleService, value string) bool {
	count := service.CheckDuplicateNameRole(value)

	return count == 0
}
