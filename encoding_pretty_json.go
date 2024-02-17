package prettygo

import (
	"encoding/json"
)

type JSONPrettyCoding struct {
}

func (coder *JSONPrettyCoding) Encode(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "   ")
}

func (coder *JSONPrettyCoding) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
