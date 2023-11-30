package pkg

import "encoding/json"

func JsonMarshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	return data, err
}
