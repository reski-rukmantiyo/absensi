package tools

import (
	"encoding/json"
	"log"
)

func ToJSON(obj interface{}) (jsonString []byte, err error) {
	jsonString, err = json.Marshal(obj)
	return
}

func FromJSON(jsonString string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		log.Println("ABS24-Error parsing JSON. " + err.Error())
	}
}
