package prettygo

import (
	"errors"
	"fmt"
	"strings"
)

type Coding interface {
	Encode(v interface{}) ([]byte, error)
	Decode(d []byte, v interface{}) error
}

func CodingFromType(codingType string) (Coding, error) {
	k := strings.ToLower(codingType)
	switch k {
	case TOML:
		return &TomlCoding{}, nil
	case JSON:
		return &JSONCoding{}, nil
	case "jsonpretty":
		fallthrough
	case "json_pretty":
		fallthrough
	case "pretty_json":
		fallthrough
	case "prettyjson":
		return &JSONPrettyCoding{}, nil
	case "yml":
		fallthrough
	case YAML:
		return &YamlCoding{}, nil
	case MSGPACK:
		return &MsgPackCoding{}, nil
	default:
		return nil, fmt.Errorf("not support type %s", codingType)
	}
}

func Encode(codingType string, v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("error invalid param")
	}
	coder, err := CodingFromType(codingType)
	if err != nil {
		return nil, err
	}
	return coder.Encode(v)
}

func Decode(codingType string, b []byte, v interface{}) error {
	if v == nil {
		return errors.New("error invalid param v")
	}

	if len(b) == 0 {
		return errors.New("error invalid param bytes")
	}

	coder, err := CodingFromType(codingType)
	if err != nil {
		return err
	}
	return coder.Decode(b, v)
}
