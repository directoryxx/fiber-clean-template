package infrastructure

import (
	"os"

	"github.com/streadway/amqp"
)

func RabbitConn() (ch *amqp.Connection, err error) {
	connRabbitMQ, err := amqp.Dial("amqp://" + os.Getenv("RABBITMQ_USER") + ":" + os.Getenv("RABBITMQ_PASS") + "@" + os.Getenv("RABBITMQ_HOST") + ":" + os.Getenv("RABBITMQ_PORT") + "/")
	// if err != nil {
	// 	panic(err)
	// }

	return connRabbitMQ, err

	// return ch, errCh
}
