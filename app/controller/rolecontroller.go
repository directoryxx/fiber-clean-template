package controller

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/middleware"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var pageRole string = "role"

// A UserController belong to the interface layer.
type RoleController struct {
	Roleservice service.RoleService
	Logger      interfaces.Logger
	Fiber       *fiber.App
	Enforcer    *casbin.Enforcer
}

func NewRoleController(logger interfaces.Logger, fiber *fiber.App, enforcer *casbin.Enforcer) *RoleController {
	return &RoleController{
		Roleservice: service.RoleService{
			RoleRepository: repository.RoleRepository{
				// SQLHandler: sqlHandler,
				Ctx: context.Background(),
			},
		},
		Logger:   logger,
		Fiber:    fiber,
		Enforcer: enforcer,
	}
}

func (controller RoleController) RoleRouter() {
	// controller.Fiber.Group(pageRole)
	controller.Fiber.Get("/role", middleware.CheckPermission(controller.Enforcer, pageRole), controller.getAll)
	controller.Fiber.Post("/role", middleware.CheckPermission(controller.Enforcer, pageRole), controller.createRole)
	controller.Fiber.Get("/role/:id", middleware.CheckPermission(controller.Enforcer, pageRole), controller.getRole)
	controller.Fiber.Put("/role/:id", middleware.CheckPermission(controller.Enforcer, pageRole), controller.updateRole)
	controller.Fiber.Delete("/role/:id", middleware.CheckPermission(controller.Enforcer, pageRole), controller.deleteRole)
}

// List Role
// @Summary List Role
// @Description List Role
// @Tags Role
// @Param Authorization header string true "With the bearer started"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /role [get]
func (controller RoleController) getAll(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	page := c.Query("page")
	limit := c.Query("limit")

	dataRole, errDataRole := controller.Roleservice.GetAll(page, limit)

	if errDataRole == 1 {
		c.Status(422)
		return c.JSON(response.ErrorResponse{
			Success: false,
			Message: "Pastikan parameter sudah benar",
		})
	}

	// token, err := jwt.ExtractTokenMetadata(c)
	// if err != nil {
	// 	controller.Logger.LogError("%s", err)
	// }

	// // res, errGet := controller.Userservice.CurrentUser(token.UserId)

	// if errGet != nil {
	// 	controller.Logger.LogError("%s", errGet)
	// }

	return c.JSON(&response.SuccessResponse{
		Success: true,
		Data:    dataRole,
		Message: "Role berhasil ditampilkan",
	})

}

// Create Role
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept application/json
// @Param Authorization header string true "With the bearer started"
// @Param role body rules.RoleValidation true "role"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /role [post]
func (controller RoleController) createRole(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	var role *domain.Role
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

// Update Role
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept application/json
// @Param role body rules.RoleValidation true "role"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /role/:id [put]
func (controller RoleController) updateRole(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	var role *domain.Role
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

// Get Role
// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /role/:id [get]
func (controller RoleController) getRole(c *fiber.Ctx) error {
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

// Delete Role
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /role/:id [delete]
func (controller RoleController) deleteRole(c *fiber.Ctx) error {
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

func roleValidation(initval *validator.Validate, service service.RoleService) {
	initval.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		return uniqueRole(service, fl.Field().String())
	})
}

func uniqueRole(service service.RoleService, value string) bool {
	count := service.CheckDuplicateNameRole(value)

	return count == 0
}
