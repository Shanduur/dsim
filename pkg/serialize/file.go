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

// BinaryFileToProtobuf deserializes the binary file into the gRPC message
func BinaryFileToProtobuf(filename string, message proto.Message) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read binary file: %v", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("unable to unmarshall binary data into proto message: %v", err)
	}

	return nil
}

// ProtobufToJSONFile saves the proto message into the JSON file
func ProtobufToJSONFile(message proto.Message, filename string) (err error) {
	json, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message into json: %v", err)
	}

	err = ioutil.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		return fmt.Errorf("unable to write json string into file: %v", err)
	}

	return nil
}
