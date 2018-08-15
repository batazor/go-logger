package main

import (
	"encoding/json"
	"github.com/batazor/go-logger/tests/amqp_generator_packets/amqp"
	"github.com/batazor/go-logger/utils"
	"time"
)

func main() {
	packetCh := make(chan interface{}, 1)
	var task = func() {
		time.Sleep(time.Millisecond * 100)
		packet, _ := utils.GetRandomPacket()
		packetCh <- packet
	}
	go task()

	for {
		select {
		case res := <-packetCh:
			json, _ := json.Marshal(res)
			//logrus.Info("json", string(json))

			for i := 0; i < 100; i++ {
				amqp.Publish([]byte(json))
			}

			go task()
		}
	}
}
