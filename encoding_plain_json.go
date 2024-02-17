package prettygo

import (
	"encoding/json"
)

type JSONCoding struct {
}

func (coder *JSONCoding) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (coder *JSONCoding) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
