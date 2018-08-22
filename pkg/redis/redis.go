package redis

import (
	"github.com/batazor/go-logger/utils"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	err error
	log = logrus.New()

	REDIS_URL = utils.Getenv("REDIS_URL", "localhost:6379")
	DB_ID     = utils.Getenv("DB_ID", "_oid")

	// Channel
	PacketCh = make(chan []byte)

	client *redis.Client
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Listen() {
	client = redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error("Redis connect error: ", err)
		return
	}

	log.Info("Run Redis")

	go func() {
		for {
			select {
			case packet := <-PacketCh:
				Insert(packet)
			}
		}
	}()
}
