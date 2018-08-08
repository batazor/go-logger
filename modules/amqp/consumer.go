package amqp

import (
	"errors"
	"github.com/batazor/go-logger/utils"
	"github.com/streadway/amqp"
	"strings"
)

func NewConsumer(uri, changes, exchangeType, queuName, key, ctag string, packetCh chan []byte) (Consumer, error) {
	var err error

	AMQP := &Consumer{
		uri:          uri,
		changes:      changes,
		exchangeType: exchangeType,
		conn:         nil,
		channel:      nil,
		tag:          ctag,
		done:         make(chan error),
	}

	AMQP.conn, err = amqp.Dial(uri)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	go func() {
		err := <-AMQP.conn.NotifyClose(make(chan *amqp.Error))
		utils.FailOnError(err, "Notify close:")
	}()

	AMQP.channel, err = AMQP.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	exchangeList := strings.Split(changes, ",")
	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = AMQP.channel.ExchangeDeclare(
			name,
			exchangeType,
			false,
			false,
			false,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to declare the Exchange")
	}

	queue, err := AMQP.channel.QueueDeclare(
		queuName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare a queue")

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = AMQP.channel.QueueBind(
			queue.Name,
			key,
			name,
			false,
			nil,
		)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	deliveries, err := AMQP.channel.Consume(
		queue.Name,
		AMQP.tag,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go handle(deliveries, AMQP.done, packetCh)

	return *AMQP, nil
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

func (c *Consumer) Connect() error {
	var err error

	c.conn, err = amqp.Dial(c.uri)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	//defer AMQP.conn.Close()

	go func() {
		// Waits here for the channel to be closed
		err := <-c.conn.NotifyClose(make(chan *amqp.Error))
		utils.FailOnError(err, "Notify close:")

		// Let Handle know it's not time to reconnect
		c.done <- errors.New("Channel Closed")
	}()

	c.channel, err = c.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	//defer AMQP.channel.Close()

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

	return nil
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
