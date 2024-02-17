package prettygo

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type TomlCoding struct {
}

func (coder *TomlCoding) Encode(v interface{}) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	defer b.Reset()

	if err := toml.NewEncoder(b).Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (coder *TomlCoding) Decode(b []byte, v interface{}) error {
	return toml.Unmarshal(b, v)
}
