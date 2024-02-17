package prettygo

import (
	"fmt"
	"strings"
)

const (
	JSON    = "json"
	TOML    = "toml"
	YAML    = "yaml"
	MSGPACK = "msgpack"
)

func CodingTypeFromExtension(ext string) (string, error) {
	extension := strings.TrimPrefix(strings.ToLower(ext), ".")
	switch extension {
	case JSON:
		return JSON, nil
	case TOML:
		return TOML, nil
	case YAML:
		return YAML, nil
	case "yml":
		return YAML, nil
	case "mp":
		return MSGPACK, nil
	default:
		return "", fmt.Errorf("not support extension: %s", extension)
	}
}
