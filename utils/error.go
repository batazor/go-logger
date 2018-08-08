package utils

import "github.com/sirupsen/logrus"

var (
	log = logrus.New()
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Warn(msg, err)
	}
}
