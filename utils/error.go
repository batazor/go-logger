package utils

import "github.com/sirupsen/logrus"

var (
	log = logrus.New()
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}
}
