package rmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getConnection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.3:5672/")
	if err != nil {
		return nil, nil, err
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return conn, ch, nil
}

func SendMsg(msg string) error {
	conn, ch, err := getConnection()
	defer conn.Close()
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"words-in", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := msg
	err = ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [>>] Sent %s\n", body)
	return nil
}

func ReceiveMsgs() {
	//
	conn, ch, err := getConnection()
	defer conn.Close()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	que, err := ch.QueueDeclare(
		//"queue.one", // name
		"queue.one", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		que.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			log.Printf("[<<] Received a message :: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
