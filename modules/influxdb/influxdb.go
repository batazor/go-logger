package influxdb

import (
	"encoding/json"
	"github.com/batazor/go-logger/utils"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	log = logrus.New()

	DB_URL      = utils.Getenv("DB_URL", "http://localhost:8086")
	DB_NAME     = utils.Getenv("DB_NAME", "telemetry")
	DB_USERNAME = utils.Getenv("DB_USERNAME", "telemetry")
	DB_PASSWORD = utils.Getenv("DB_PASSWORD", "telemetry")
	DB_ID       = utils.Getenv("DB_ID", "_oid")

	CLIENT client.Client
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Connect(packetCh chan []byte) {
	// Create a new HTTPClient
	CLIENT, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     DB_URL,
		Username: DB_USERNAME,
		Password: DB_PASSWORD,
	})
	utils.FailOnError(err, "Error create a new HTTPClient")
	defer CLIENT.Close()

	var wait = func() {
		for {
			select {
			case packet := <-packetCh:
				// Parse
				log.Info("JSON: ", string(packet))
				fields := map[string]interface{}{}
				err := json.Unmarshal(packet, &fields)
				utils.FailOnError(err, "Error parse packet")

				// Create a new point batch
				bp, err := client.NewBatchPoints(client.BatchPointsConfig{
					Database:  DB_NAME,
					Precision: "s",
				})
				utils.FailOnError(err, "Error create a new point batch")

				// Create a point and add to batch
				tags := map[string]string{"telemetry": "raw"}

				pt, err := client.NewPoint(fields[DB_ID].(string), tags, fields, time.Now())
				if err != nil {
					log.Fatal(err)
				}
				bp.AddPoint(pt)

				// Write the batch
				if err := CLIENT.Write(bp); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Warn("Problem in InfluxDB", r)

			go wait()
		}
	}()
}
