package controller

import (
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
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
	conn, _ := infrastructure.RabbitConn()
	defer conn.Close()
	ch, errCh := conn.Channel()
	if errCh != nil {
		panic(errCh)
	}

	query := c.Query("msg")

	_, err := ch.QueueDeclare(
		"TestQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	// Handle any errors if we were unable to create the queue.
	if err != nil {
		panic(err)
	}

	// Attempt to publish a message to the queue.
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(query),
		},
	)
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	return c.JSON(&fiber.Map{
		"success": "ook",
	})
}
