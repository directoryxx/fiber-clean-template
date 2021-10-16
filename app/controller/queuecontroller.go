package controller

import (
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type QueueController struct {
	Logger interfaces.Logger
}

func NewQueueController(logger interfaces.Logger) *QueueController {
	return &QueueController{
		Logger: logger,
	}
}

func (controller *QueueController) TestQueue(c *fiber.Ctx) error {
	query := c.Query("msg")

	infrastructure.SendQueue(query, "TestQueue")

	return c.JSON(&fiber.Map{
		"success": "ook",
	})
}
