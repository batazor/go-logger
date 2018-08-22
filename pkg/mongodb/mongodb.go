package mongodb_new

import (
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	ID   string
	Item map[string]interface{}
}

var (
	log       = logrus.New()
	MONGO_URL = utils.Getenv("MONGO_URL", "mongodb://localhost:27017")

	// Channel
	PacketCh = make(chan Document)
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Connect() {
	client, err := mgo.Dial(MONGO_URL)
	if err != nil {
		log.Error("Error uri for MongoDB: ", err)
	}

	log.Info("Run MongoDB")

	go func() {
		for {
			select {
			case packet := <-PacketCh:
				go func() {
					id := bson.NewObjectId()
					packet.Item["_id"] = id
					err := client.DB("go-loger").C(packet.ID).Insert(packet.Item)

					if err != nil {
						log.Error("Insert mongodb error: ", err)
					}
				}()
			}
		}
	}()
}
