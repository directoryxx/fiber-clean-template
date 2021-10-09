package controller

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/middleware"
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var pagePermission string = "permission"

// A UserController belong to the interface layer.
type PermissionController struct {
	Enforcer    *casbin.Enforcer
	Logger      interfaces.Logger
	Fiber       *fiber.App
	RoleService *service.RoleService
}

func NewPermissionController(logger interfaces.Logger, enforcer *casbin.Enforcer, app *fiber.App) *PermissionController {
	return &PermissionController{
		Enforcer: enforcer,
		Logger:   logger,
		Fiber:    app,
		RoleService: &service.RoleService{
			RoleRepository: repository.RoleRepository{
				Ctx: context.Background(),
			},
		},
	}
}

func (controller PermissionController) PermissionRouter() {
	// controller.Fiber.Group(pagePermission)
	controller.Fiber.Get("/permission/:id", middleware.CheckPermission(controller.Enforcer, pagePermission), controller.getListPermission)
	controller.Fiber.Post("/permission/:id", middleware.CheckPermission(controller.Enforcer, pagePermission), controller.updatePermission)
}

// List Permission
// @Summary List Permission
// @Description List Permission by role
// @Tags Permission
// @Param Authorization header string true "With the bearer started"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /permission/:id [get]
func (controller PermissionController) getListPermission(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	id, err := c.ParamsInt("id")

	roleData := controller.RoleService.GetById(id)

	if roleData == nil {
		c.Status(404)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Data tidak ditemukan",
		})
	}

	getPolicy := controller.Enforcer.GetFilteredPolicy(0, roleData.Name)

	if err != nil {
		c.Status(422)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Silahkan periksa kembali",
		})
	}

	return c.JSON(&response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data",
		Data:    getPolicy,
	})

}

// Update Permission
// @Summary Update Permission
// @Description Update Permission by role
// @Tags Permission
// @Accept application/json
// @Param Authorization header string true "With the bearer started"
// @Param permission body rules.PermissionUpdate true "Permission"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /permission/:id [post]
func (controller PermissionController) updatePermission(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	var permission rules.PermissionUpdate
	err := c.BodyParser(&permission)

	if err != nil {
		c.Status(422)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Silahkan periksa kembali",
		})
	}

	id, err := c.ParamsInt("id")

	roleData := controller.RoleService.GetById(id)

	if roleData == nil {
		c.Status(404)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Data tidak ditemukan",
		})
	}

	initval = validator.New()
	roleValidation(initval, *controller.RoleService)
	errVal := initval.Struct(permission)

	if errVal != nil {
		message := make(map[string]string)
		message["permission"] = "Pastikan kembali data yang dikirim"
		errorResponse := validation.ValidateRequest(errVal, message)
		return c.JSON(errorResponse)
	}

	controller.Enforcer.RemoveFilteredPolicy(0, roleData.Name)

	for _, s := range permission.Permission {
		controller.Enforcer.AddPolicy(roleData.Name, s.Page, s.Resource)
	}

	getPolicy := controller.Enforcer.GetFilteredPolicy(0, roleData.Name)

	if err != nil {
		c.Status(422)
		return c.JSON(&response.ErrorResponse{
			Success: false,
			Message: "Silahkan periksa kembali",
		})
	}

	return c.JSON(&response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data",
		Data:    getPolicy,
	})

}
