package amqp

import "github.com/streadway/amqp"

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}
