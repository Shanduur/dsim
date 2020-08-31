package serialize

import (
	"os"
	"testing"

	"github.com/Sheerley/pluggabl/pkg/transfer"
)

func TestProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	name := "file.tmp.bin"

	req := transfer.NewRequest(10)
	err := ProtobufToBinaryFile(req, name)
	if err != nil {
		t.Errorf("Error saving protobuf message to binary file: %v", err)
	}

	err = os.Remove(name)
	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}
}
