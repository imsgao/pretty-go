package prettygo

import (
	"github.com/vmihailenco/msgpack"
)

type MsgPackCoding struct {
}

func (coder *MsgPackCoding) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (coder *MsgPackCoding) Decode(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}
