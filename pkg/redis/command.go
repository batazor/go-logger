package redis

import (
	"encoding/json"
	"github.com/batazor/go-logger/pkg/mongodb"
	"github.com/imdario/mergo"
)

func Insert(packetNew []byte) bool {
	fieldsNew, _ := parseJSON(packetNew)

	ID := fieldsNew[DB_ID].(string)

	res := client.Get(ID)
	if res.Err() == nil {
		d, _ := res.Bytes()
		fieldsOld, _ := parseJSON(d)

		if err := mergo.Merge(&fieldsOld, fieldsNew); err != nil {
			log.Error("Error merge new and old packets", err)
			return false
		}
	}

	data, er := json.Marshal(fieldsNew)
	if er != nil {
		log.Error("Error parse JSON: ", er)
		return false
	}

	// Insert to Redis
	r := client.Set(ID, data, 0)
	if r.Err() != nil {
		log.Error("Redis SET error: ", r.Err(), " args: ", r.Args())
		return false
	}

	// Insert to MongoDB
	mongodb_new.PacketCh <- mongodb_new.Document{
		ID:   ID,
		Item: fieldsNew,
	}

	return true
}
