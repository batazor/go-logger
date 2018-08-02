package main

import (
	json2 "encoding/json"
	"github.com/batazor/go-logger/tests/bot/amqp"
	"github.com/bxcodec/faker"
	"github.com/sirupsen/logrus"
	"time"
)

type State struct {
	Latitude float32 `faker:"lat" json:"lat"`
	Long     float32 `faker:"long" json:"lon"`
	Time     string  `faker:"time" json:"time"`
}

type Packet struct {
	Oid       string `faker:"username" json:"_oid"`
	UserName  string `faker:"username" json:"username"`
	UnixTime  int64  `faker:"unix_time" json:"unixtime"`
	Date      string `faker:"date" json:"date"`
	MonthName string `faker:"month_name" json:"monthName"`
	Year      string `faker:"year" json:"year"`
	DayOfWeek string `faker:"day_of_week" json:"dayOfWeek"`
	Timestamp string `faker:"timestamp" json:"timestamp"`
	TimeZone  string `faker:"timezone"  json:"timezone"`
	IPV4      string `faker:"ipv4" json:"IPv4"`
	State     State  `json:"state"`
}

func getRandomPacket() (interface{}, error) {
	a := Packet{}
	err := faker.FakeData(&a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func main() {
	packetCh := make(chan interface{}, 1)
	go func() {
		time.Sleep(time.Millisecond * 10)
		packet, _ := getRandomPacket()
		packetCh <- packet
	}()

	select {
	case res := <-packetCh:
		json, _ := json2.Marshal(res)
		logrus.Info("json", string(json))
		amqp.Publish([]byte(json))
	}
}
