package redis

import (
	"encoding/json"
	"github.com/jeremywohl/flatten"
)

func parseJSON(packet []byte) (map[string]interface{}, error) {
	flat, err := flatten.FlattenString(string(packet), "", flatten.DotStyle)
	if err != nil {
		log.Error("Parse JSON error: ", err)
		return nil, err
	}

	fields := map[string]interface{}{}
	err = json.Unmarshal([]byte(flat), &fields)
	if err != nil {
		log.Warn("Error parse packet: ", err)
		return nil, err
	}

	return fields, nil
}
