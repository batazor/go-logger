package influxdb

import (
	probe "github.com/batazor/go-logger/pkg/healthcheck"
	"github.com/batazor/go-logger/utils"
	"github.com/heptiolabs/healthcheck"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	err error
	log = logrus.New()

	DB_URL      = utils.Getenv("DB_URL", "http://localhost:8086")
	DB_NAME     = utils.Getenv("DB_NAME", "telemetry")
	DB_USERNAME = utils.Getenv("DB_USERNAME", "telemetry")
	DB_PASSWORD = utils.Getenv("DB_PASSWORD", "telemetry")
	DB_ID       = utils.Getenv("DB_ID", "_oid")

	SESSION client.Client

	// Channel
	PacketCh = make(chan []byte)
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

// Create a new HTTPClient
func influxDBClient() (client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     DB_URL,
		Username: DB_USERNAME,
		Password: DB_PASSWORD,
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Connect() {
	SESSION, err = influxDBClient()
	if err != nil {
		log.Warn("Error create a new HTTPClient: ", err)
	}
	log.Info("Run InfluxDB")

	// Health check
	probe.Health.AddReadinessCheck(
		"influxdb",
		healthcheck.Timeout(func() error { return err }, time.Second*10))

	go func() {
		for {
			select {
			case packet := <-PacketCh:
				InsertJSON(string(packet))
			}
		}
	}()

	// TEST
	response := getCountPointByMeasurements()

	for _, v := range response.Results[0].Series {
		r := StateRequest{
			measurement: `"` + v.Name + `"`,
			function:    "last",
			fields:      "*",
			where:       "time > now() - 7d",
		}
		d := GetState(r)
		log.Info("D: ", d.Results[0].Series)

		log.Info("T: ", v.Name, " : ", v.Values[0][1])
	}
}
