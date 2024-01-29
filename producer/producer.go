package producer

import (
	"context"
	"log"
	"rabbitgo/connection"
	corelib "rabbitgo/coreLib"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Producer(ctx context.Context, msg string) {

	err := connection.Channel.PublishWithContext(ctx,
		"fanout-example", // exchange
		"",               // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		}, //message
	)
	corelib.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", msg)

}
