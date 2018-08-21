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

func getSizeDB() *client.Response {
	q := client.Query{
		Command:  `SELECT last("diskBytes") FROM "_internal"."monitor"."shard" GROUP BY "database"`,
		Database: "_internal",
	}

	log.Info("REQUEST: ", q.Command)

	response, err := SESSION.Query(q)
	if err != nil {
		log.Info("Error: ", err)
		return nil
	}

	if response.Error() != nil {
		log.Info("Response error: ", response.Error())
		return nil
	}

	if response.Err != "" {
		log.Info("Serie error: ", response.Err)
		return nil
	}

	return response
}
