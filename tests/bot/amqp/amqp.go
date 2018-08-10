package amqp

import (
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strings"
	"time"
)

var (
	log = logrus.New()

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://pb:pb@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "go-logger-packets")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")

	AMQP_CH amqp.Channel
	AMQP_Q  amqp.Queue
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Publish(message []byte) error {
	AMQP_CONN, err := amqp.Dial(AMQP_API)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer AMQP_CONN.Close()

	AMQP_CH, err := AMQP_CONN.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer AMQP_CH.Close()

	exchangeList := strings.Split(AMQP_EXCHANGE_LIST, ",")
	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = AMQP_CH.ExchangeDeclare(
			name,
			"headers",
			false,
			false,
			false,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to declare the Exchange")
	}

	AMQP_Q, err := AMQP_CH.QueueDeclare(
		AMQP_NAME_QUEUE,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare a queue")

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")

		err = AMQP_CH.QueueBind(
			AMQP_Q.Name,
			"",
			name,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to bind a queue")

		err = AMQP_CH.Publish(
			"",
			AMQP_Q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:  "application/json",
				DeliveryMode: amqp.Transient,
				Body:         message,
				Timestamp:    time.Now(),
			},
		)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	return nil
}
