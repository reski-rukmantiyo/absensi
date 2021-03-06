package tools

import "encoding/json"

func ToJSON(obj interface{}) (jsonString []byte, err error) {
	jsonString, err = json.Marshal(obj)
	return
}

func FromJSON(jsonString string, obj interface{}) interface{} {
	objReturn := json.Unmarshal([]byte(jsonString), obj)
	return objReturn
}
