package infrastructure

import (
	"fmt"
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

func SendQueue(data interface{}, queuename string) {
	// fmt.Println("data : ", data)
	str := fmt.Sprintf("%v", data)
	conn, _ := RabbitConn()
	defer conn.Close()
	ch, errCh := conn.Channel()
	if errCh != nil {
		panic(errCh)
	}
	_, err := ch.QueueDeclare(
		queuename,
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
		queuename,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(str),
		},
	)
	if err != nil {
		panic(err)
	}

	defer ch.Close()
}
