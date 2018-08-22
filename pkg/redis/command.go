package redis

import (
	"encoding/json"
	"github.com/imdario/mergo"
)

func Insert(packetNew []byte) []byte {
	fieldsNew, _ := parseJSON(packetNew)

	ID := fieldsNew[DB_ID].(string)

	res := client.Get(ID)
	if res.Err() == nil {
		d, _ := res.Bytes()
		fieldsOld, _ := parseJSON(d)

		if err := mergo.Merge(&fieldsOld, fieldsNew); err != nil {
			log.Error("Error merge new and old packets", err)
		}
	}

	data, er := json.Marshal(fieldsNew)
	if er != nil {
		log.Error("Error parse JSON: ", er)
	}

	r := client.Set(ID, data, 0)
	if r.Err() != nil {
		log.Error("Redis SET error: ", r.Err(), " args: ", r.Args())
	}

	return []byte(packetNew)
}
