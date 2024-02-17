package prettygo

func EncodeJSON(v interface{}) ([]byte, error) {
	return Encode(JSON, v)
}

func EncodePrettyJSON(v interface{}) ([]byte, error) {
	return Encode("prettyjson", v)
}

func DecodeJSON(b []byte, v interface{}) error {
	return Decode(JSON, b, v)
}

func EncodeTOML(v interface{}) ([]byte, error) {
	return Encode(TOML, v)
}

func DecodeTOML(b []byte, v interface{}) error {
	return Decode(TOML, b, v)
}

func EncodeYAML(v interface{}) ([]byte, error) {
	return Encode(YAML, v)
}

func DecodeYAML(b []byte, v interface{}) error {
	return Decode(YAML, b, v)
}

func EncodeMsgPack(v interface{}) ([]byte, error) {
	return Encode(MSGPACK, v)
}

func DecodeMsgPack(b []byte, v interface{}) error {
	return Decode(MSGPACK, b, v)
}
