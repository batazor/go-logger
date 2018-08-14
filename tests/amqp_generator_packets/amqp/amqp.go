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

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://telemetry:telemetry@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "go-logger-packets")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")
	AMQP_EXCHANGE_TYPE = utils.Getenv("AMQP_EXCHANGE_TYPE", "headers")
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
	if err != nil {
		log.Info("Failed to connect to RabbitMQ: ", err)
	}
	defer AMQP_CONN.Close()

	AMQP_CH, err := AMQP_CONN.Channel()
	if err != nil {
		log.Info("Failed to open a channel: ", err)
	}
	defer AMQP_CH.Close()

	exchangeList := strings.Split(AMQP_EXCHANGE_LIST, ",")
	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = AMQP_CH.ExchangeDeclare(
			name,
			AMQP_EXCHANGE_TYPE,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Info("Failed to declare the Exchange: ", err)
		}
	}

	AMQP_Q, err := AMQP_CH.QueueDeclare(
		AMQP_NAME_QUEUE,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Info("Failed to declare a queue: ", err)
	}

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")

		err = AMQP_CH.QueueBind(
			AMQP_Q.Name,
			"",
			name,
			false,
			nil,
		)
		if err != nil {
			log.Info("Failed to bind a queue: ", err)
		}

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
		if err != nil {
			log.Info("Failed to bind a queue: ", err)
		}
	}

	return nil
}
