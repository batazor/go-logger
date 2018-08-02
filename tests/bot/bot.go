package main

import (
	json2 "encoding/json"
	"github.com/batazor/go-logger/tests/bot/amqp"
	"github.com/bxcodec/faker"
	"time"
)

type State struct {
	Latitude float32 `faker:"lat" json:"lat"`
	Long     float32 `faker:"long" json:"lon"`
	Time     string  `faker:"time" json:"time"`
}

type Packet struct {
	Oid       string `faker:"month_name" json:"_oid"`
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
	Bool      bool
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
	var task = func() {
		time.Sleep(time.Millisecond * 100)
		packet, _ := getRandomPacket()
		packetCh <- packet
	}
	go task()

	for {
		select {
		case res := <-packetCh:
			json, _ := json2.Marshal(res)
			//logrus.Info("json", string(json))

			for i := 0; i < 100; i++ {
				amqp.Publish([]byte(json))
			}

			go task()
		}
	}
}
