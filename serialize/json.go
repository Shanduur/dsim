package serialize

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON serializes protobuf message into JSON
func ProtobufToJSON(message proto.Message) (string, error) {
	b := protojson.MarshalOptions{
		Multiline:       true,
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	return b.Format(message), nil
}
