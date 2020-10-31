package utils

import "github.com/golang/protobuf/proto"

func ProtoMarshal(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func ProtoUnmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}
