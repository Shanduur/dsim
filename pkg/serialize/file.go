package serialize

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// ProtobufToBinaryFile serializes the gRPC message into the file
func ProtobufToBinaryFile(message proto.Message, filename string) (err error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message into binary: %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write proto message data into file: %v", err)
	}

	return nil
}
