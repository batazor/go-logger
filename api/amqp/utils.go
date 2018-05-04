package amqp

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}
}