package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/directoryxx/fiber-clean-template/app/utils/session"
	"github.com/gofiber/fiber/v2"
)

// TODO : Refactor This
func CheckPermission(enforcer *casbin.Enforcer, page string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := session.GetAuth()

		role := auth.Role

		ok, err := enforcer.Enforce(role.Name, page, c.Method())

		okManage, _ := enforcer.Enforce(role.Name, page, "manage")

		if err != nil {
			c.Status(500)
			return c.JSON(response.ErrorResponse{
				Success: false,
				Message: err.Error(),
				// Error:   err,
			})
		}

		if okManage {
			return c.Next()
		}

		if !ok {
			// errorForbidden := errors.New("unauthorized access")
			c.Status(403)
			return c.JSON(response.ErrorResponse{
				Success: false,
				Message: "Forbidden access",
				// Error:   errorForbidden,
			})
		}

		return c.Next()
	}
}
