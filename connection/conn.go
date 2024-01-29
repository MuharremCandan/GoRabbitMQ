package connection

import (
	corelib "rabbitgo/coreLib"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Channel *amqp.Channel
)

func InitConnRabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	corelib.FailOnError(err, "Failed to connect to RabbitMQ")

	channel, err := conn.Channel()
	corelib.FailOnError(err, "Failed to open a channel")

	Channel = channel

	exchange()

}

func exchange() {
	if err := Channel.ExchangeDeclare("fanout-example", "fanout", true, false, false, false, nil); err != nil {
		corelib.FailOnError(err, "Failed to declare a example")
	}
}
