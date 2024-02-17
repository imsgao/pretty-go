package prettygo

import (
	"github.com/ghodss/yaml"
)

type YamlCoding struct {
}

func (coder *YamlCoding) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (coder *YamlCoding) Decode(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}
