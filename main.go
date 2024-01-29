package main

import (
	"context"
	queue "rabbitgo/Queue"
	"rabbitgo/connection"
	"rabbitgo/consumer"
	corelib "rabbitgo/coreLib"
	"rabbitgo/producer"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {

	connection.InitConnRabbit()

	queue1 := queue.InitQueue("TestQueue") // give a name to your queue

	queue2 := queue.InitQueue("TestQueue2")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	my_exchange := "fanout-example"

	QueueBind(queue1, my_exchange)

	QueueBind(queue2, my_exchange)

	msg := "My message to RabbitMq"

	producer.Producer(ctx, msg)

	consumer.Consumer(queue1)

}

func QueueBind(queue amqp091.Queue, excahnge string) {
	if err := connection.Channel.QueueBind(queue.Name, "", excahnge, false, nil); err != nil {
		corelib.FailOnError(err, "Failed to bind to %s"+queue.Name)
	}
}
