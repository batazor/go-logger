package main

import (
	"fmt"
	"github.com/batazor/go-logger/modules/amqp"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	log = logrus.New()

	packetCh = make(chan []byte)
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func main() {
	go influxdb.Connect(packetCh)
	go amqp.Listen(packetCh)

	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Info("Listen HTTP 8080")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
