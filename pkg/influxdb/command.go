package influxdb

import (
	"encoding/json"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jeremywohl/flatten"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Query(DataBase string) []byte {
	q := client.Query{
		Command:  `SELECT LAST("year") from "telemetry"."autogen"."` + DataBase + `"`,
		Database: DB_NAME,
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

	result := response.Results[0]
	if result.Err != "" {
		log.Info("Serie error: ", result.Err)
		return nil
	}

	// GO struct to JSON schema
	b, err := json.Marshal(result.Series)
	if err != nil {
		log.Error("JSON marshal error: ", result.Err)
	}

	log.Info("RES: ", string(b))

	return b
}

func Insert(fields map[string]interface{}) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB_NAME,
		Precision: "s",
	})
	if err != nil {
		log.Warn("Error create a new point batch: ", err)
		return err
	}

	// Create a point and add to batch
	tags := map[string]string{"telemetry": "raw"}

	pt, err := client.NewPoint(fields[DB_ID].(string), tags, fields, time.Now())
	if err != nil {
		log.Error("Error create new point: ", err)
		return err
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := SESSION.Write(bp); err != nil {
		log.Error("Error write new point: ", err)
		return err
	}

	return nil
}

func InsertJSON(packet string) bool {
	// Parse
	// Nested JSON to flat JSON
	flat, err := flatten.FlattenString(string(packet), "", 0)
	if err != nil {
		log.Warn("Error convert nested JSON to flat JSON: ", err)
	}

	fields := map[string]interface{}{}
	err = json.Unmarshal([]byte(flat), &fields)
	if err != nil {
		log.Warn("Error parse packet: ", err)
	}

	if err = Insert(fields); err != nil {
		return false
	}

	return true
}

func GetState(request StateRequest) *client.Response {
	q := client.Query{
		Command:  `SELECT ` + request.function + `("` + request.fields + `") FROM "` + DB_NAME + `"."autogen".` + request.measurement,
		Database: DB_NAME,
	}

	if request.where != "" {
		q.Command = q.Command + " WHERE " + request.where
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
