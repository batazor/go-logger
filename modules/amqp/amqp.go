package amqp

import (
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	log = logrus.New()

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://telemetry:telemetry@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "go-logger-packets")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")

	AMQP_CH amqp.Channel
	AMQP_Q  amqp.Queue

	forever      = make(chan bool)
	gracefulStop = make(chan os.Signal)
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Gracefully stop application
	signal.Notify(gracefulStop, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-gracefulStop:
				exchangeList := strings.Split(AMQP_EXCHANGE_LIST, ",")
				for _, echangeName := range exchangeList {
					name := strings.Trim(echangeName, " ")
					err := AMQP_CH.QueueUnbind(
						AMQP_Q.Name,
						name,
						"",
						nil,
					)
					utils.FailOnError(err, "Failed to unbind a queue")
				}
			}
		}
	}()
}

func Listen(packetCh chan []byte) {
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
	}

	msgs, err := AMQP_CH.Consume(
		AMQP_Q.Name,
		"go-logger",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			packetCh <- d.Body
		}
	}()

	log.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
