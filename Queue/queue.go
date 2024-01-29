package queue

import (
	"rabbitgo/connection"
	corelib "rabbitgo/coreLib"

	"github.com/rabbitmq/amqp091-go"
)

func InitQueue(queuename string) amqp091.Queue {

	queue, err := connection.Channel.QueueDeclare(
		queuename, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	corelib.FailOnError(err, "Failed to declare a queue")

	return queue
}
