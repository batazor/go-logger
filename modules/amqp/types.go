package amqp

import "github.com/streadway/amqp"

type Consumer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	tag          string
	uri          string
	changes      string
	exchangeType string
	done         chan error
}
