package proxy

import (
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

// Codec returns a proxying grpc.Codec with the default protobuf codec as parent.
//
// See CodecWithParent.
func Codec(c encoding.Codec) encoding.Codec {
	return &protoCodec{c}
}

// protoCodec is a Codec implementation with protobuf. It is the default rawCodec for gRPC.
type protoCodec struct {
	c encoding.Codec
}

func (c *protoCodec) Marshal(v any) ([]byte, error) {
	return c.c.Marshal(v.(proto.Message))
}

func (c *protoCodec) Unmarshal(data []byte, v any) error {
	return c.c.Unmarshal(data, v.(proto.Message))
}

func (protoCodec) Name() string {
	return "proto"
}
