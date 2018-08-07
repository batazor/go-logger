package amqp

import (
	"github.com/batazor/go-logger/utils"
	"github.com/streadway/amqp"
	"strings"
)

func NewConsumer(url, changes, exchangeType, queuName, key, ctag string, packetCh chan []byte) (Consumer, error) {
	var err error

	AMQP := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),
	}

	AMQP.conn, err = amqp.Dial(url)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	//defer AMQP.conn.Close()

	go func() {
		err := <-AMQP.conn.NotifyClose(make(chan *amqp.Error))
		utils.FailOnError(err, "Notify close:")
	}()

	AMQP.channel, err = AMQP.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	//defer AMQP.channel.Close()

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

	//log.Info("declared Queue (", queue.Name, ": ", queue.Messages, " messages, ", queue.Consumers, " consumers), binding to Exchange (key: ", key, ")")

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

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
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

func handle(deliveries <-chan amqp.Delivery, done chan error, packetCh chan []byte) {
	for d := range deliveries {
		packetCh <- d.Body
		d.Ack(false)
	}
	log.Info("handle: deliveries channel closed")
	done <- nil
}
