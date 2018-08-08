package amqp

import (
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

var (
	log = logrus.New()

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://telemetry:telemetry@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "go-logger-packets")
	AMQP_BINDING_KEY   = utils.Getenv("AMQP_BINDING_KEY", "")
	AMQP_CONSUMER_TAG  = utils.Getenv("AMQP_CONSUMER_TAG", "")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")
	AMQP_EXCHANGE_TYPE = utils.Getenv("AMQP_EXCHANGE_TYPE", "headers")

	gracefulStop = make(chan os.Signal)

	CONSUMER = &Consumer{}
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Gracefully stop application
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-gracefulStop:
				// Close connect to AMQP
				if err := CONSUMER.Shutdown(); err != nil {
					log.Error("Failed shutdown AMQP")
				}
			}
		}
	}()
}

func Listen(packetCh chan []byte) {
	CONSUMER = NewConsumer(AMQP_API, AMQP_EXCHANGE_LIST, AMQP_EXCHANGE_TYPE, AMQP_NAME_QUEUE, AMQP_BINDING_KEY, AMQP_CONSUMER_TAG, packetCh)

	if err := CONSUMER.Connect(); err != nil {
		log.Warn(err)
	}

	deliveries, err := CONSUMER.AnnounceQueue()
	if err != nil {
		log.Warn(err)
	}

	CONSUMER.Handle(deliveries, handler)
}

func handler(deliveries <-chan amqp.Delivery) {
	threads := utils.MaxParallelism()

	for i := 0; i < threads; i++ {
		go func() {
			for d := range deliveries {
				//packetCh <- d.Body
				d.Ack(false)
			}
		}()
	}

	return
}
