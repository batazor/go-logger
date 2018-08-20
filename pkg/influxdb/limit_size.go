package influxdb

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func getCountPointByMeasurements() *client.Response {
	r := StateRequest{
		measurement: "/^*/",
		function:    "count",
		fields:      DB_ID,
	}

	return GetState(r)
}
