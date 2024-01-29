package consumer

import (
	"log"
	"rabbitgo/connection"
	corelib "rabbitgo/coreLib"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer(queue amqp.Queue) {
	msgs, err := connection.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	corelib.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message by consumer1: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
