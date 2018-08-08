package amqp

import (
	"github.com/batazor/go-logger/utils"
	"github.com/streadway/amqp"
	"strings"
)

func NewConsumer(uri, changes, exchangeType, queueName, bindingKey, consumerTag string, packetCh chan []byte) *Consumer {
	return &Consumer{
		uri:          uri,
		changes:      changes,
		bindingKey:   bindingKey,
		exchangeType: exchangeType,
		queueName:    queueName,
		conn:         nil,
		channel:      nil,
		consumerTag:  consumerTag,
		done:         make(chan error),
		packetCh:     packetCh,
	}

}

func (c *Consumer) Connect() error {
	var err error

	c.conn, err = amqp.Dial(c.uri)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	go func() {
		err := <-c.conn.NotifyClose(make(chan *amqp.Error))
		utils.FailOnError(err, "Notify close:")
	}()

	c.channel, err = c.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	exchangeList := strings.Split(c.changes, ",")
	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = c.channel.ExchangeDeclare(
			name,
			c.exchangeType,
			false,
			false,
			false,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to declare the Exchange")
	}

	queue, err := c.channel.QueueDeclare(
		c.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare a queue")

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = c.channel.QueueBind(
			queue.Name,
			c.bindingKey,
			name,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	deliveries, err := c.channel.Consume(
		queue.Name,
		c.consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go handle(deliveries, c.done, c.packetCh)

	return nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error, packetCh chan []byte) {
	threads := utils.MaxParallelism()

	for i := 0; i < threads; i++ {
		go func() {
			for d := range deliveries {
				packetCh <- d.Body
				d.Ack(false)
			}
		}()
	}

	log.Info("handle: deliveries channel closed")
	done <- nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Close(); err != nil {
		log.Error("Consumer cancel failed", err)
		return err
	}

	if err := c.conn.Close(); err != nil {
		log.Error("AMQP connection close error", err)
		return err
	}

	defer log.Warn("AMQP shutwodn OK")

	// wait for handle() to exit
	return <-c.done
}
