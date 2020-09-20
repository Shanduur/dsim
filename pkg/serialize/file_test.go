package serialize

import (
	"testing"

	"github.com/Sheerley/pluggabl/pkg/pb"
	"google.golang.org/protobuf/proto"

	"github.com/Sheerley/pluggabl/pkg/transfer"
)

func TestProtoAndBinaryFile(t *testing.T) {
	t.Parallel()

	name := "file.tmp.bin"

	req1 := transfer.NewDummyJob(10)
	err := ProtobufToBinaryFile(req1, name)
	if err != nil {
		t.Errorf("Error saving protobuf message to binary file: %v", err)
	}

	req2 := &pb.Job{}

	err = BinaryFileToProtobuf(name, req2)
	if err != nil {
		t.Errorf("Error reading protobuf message from binary file: %v", err)
	}

	if !proto.Equal(req1, req2) {
		t.Error("Data inside both messages is not equal")
	}
}

func TestProtoAndJSONFile(t *testing.T) {
	t.Parallel()

	name := "file.tmp.json"

	req1 := transfer.NewDummyJob(10)
	err := ProtobufToJSONFile(req1, name)
	if err != nil {
		t.Errorf("Unable to save proto message into JSON file")
	}
}
