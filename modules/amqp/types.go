package amqp

import (
	"github.com/streadway/amqp"
)

// Consumer holds all information
// about the RabbitMQ connection
// This setup does limit a consumer
// to one exchange. This should not be
// an issue. Having to connect to multiple
// exchanges means something else is
// structured improperly.
type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	done     chan error
	packetCh chan []byte

	consumerTag  string // Name that consumer identifies itself to the server with
	uri          string // uri of the rabbitmq server
	changes      string // exchange that we will bind to
	exchangeType string // topic, direct, etc...
	bindingKey   string // routing key that we are using
}
