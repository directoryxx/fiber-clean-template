package middleware

import (
	"github.com/directoryxx/fiber-clean-template/app/service"
	"github.com/directoryxx/fiber-clean-template/app/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// TODO : Refactor This
func JWTProtected(svc service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		store := session.New()
		sess, _ := store.Get(c)
		token, errExtract := jwt.ExtractTokenMetadata(c)

		if errExtract != nil {
			c.Status(401)
			return c.JSON(fiber.Map{"error": "Unauthorized access"})
		}

		userId, _ := jwt.FetchAuth(svc, token)

		if userId == 0 {
			c.Status(401)
			return c.JSON(fiber.Map{"error": "Unauthorized access"})
		}

		sess.Set("id", userId)

		return c.Next()

	}
}
