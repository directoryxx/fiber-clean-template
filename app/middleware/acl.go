package middleware

// TODO : Refactor This
// func CheckPermission(enforcer *casbin.Enforcer, service service.UserService, page string, resource string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		token, err := jwt.ExtractTokenMetadata(c)
// 		if err != nil {
// 			c.Status(403)
// 			return c.JSON(fiber.Map{"error": "Unauthorized access"})
// 		}

// 		userId, _ := jwt.FetchAuth(service, token)

// 		// roleUser, _ := service.GetRoleUser(uint(userId))

// 		ok, err := enforcer.Enforce(roleUser.RoleUser.Role.Name, page, resource)

// 		okManage, _ := enforcer.Enforce(roleUser.RoleUser.Role.Name, page, "manage")

// 		if err != nil {
// 			c.Status(500)
// 			return c.JSON(response.ErrorResponse{
// 				Success: false,
// 				Message: err.Error(),
// 				// Error:   err,
// 			})
// 		}

// 		if okManage {
// 			return c.Next()
// 		}

// 		if !ok {
// 			// errorForbidden := errors.New("unauthorized access")
// 			c.Status(403)
// 			return c.JSON(response.ErrorResponse{
// 				Success: false,
// 				Message: "Unauthorized access",
// 				// Error:   errorForbidden,
// 			})
// 		}

// 		return c.Next()
// 	}
// }
