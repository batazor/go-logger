package amqp

import (
	"github.com/sirupsen/logrus"
	"github.com/batazor/go-logger/utils"
	"github.com/streadway/amqp"
	"strings"
)

var (
	log = logrus.New()

	AMQP_API = utils.Getenv("AMQP_API", "amqp://guest:guest@localhost:5672/")
	AMQP_NAME_QUEUE = utils.Getenv("AMQP_NAME_QUEUE", "input")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")
	forever = make(chan bool)
)

func Listen() {
	conn, err := amqp.Dial(AMQP_API)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		AMQP_NAME_QUEUE,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	exchangeList := strings.Split(AMQP_EXCHANGE_LIST, ",")
	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = ch.QueueBind(
			q.Name,
			"",
			name,
			false,
			nil,
		)
		failOnError(err, "Failed to bind a queue")
	}


	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Info("Received a message: ", string(d.Body))
		}
	}()

	log.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}