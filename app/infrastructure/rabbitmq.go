package infrastructure

import "github.com/streadway/amqp"

func RabbitConn() (ch *amqp.Connection, err error) {
	connRabbitMQ, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// if err != nil {
	// 	panic(err)
	// }

	return connRabbitMQ, err

	// return ch, errCh
}
