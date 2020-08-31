package transfer

import (
	"github.com/Sheerley/pluggabl/pkg/pb"
)

// NewCredentials generates Credentials struct
func NewCredentials() *pb.Credentials {
	userCreds := &pb.Credentials{
		UserId:  "placeholder",
		UserKey: "placeholder",
	}

	return userCreds
}

// NewFileInfo generates File information
func NewFileInfo() *pb.FileInfo {
	fileInfo := &pb.FileInfo{
		FileExtension: ".txt",
		SizeType:      pb.FileInfo_kilobytes,
		Size:          2048,
	}

	return fileInfo
}

// NewFilesSlice generates slice containing informations about transfered files
func NewFilesSlice(number uint32) []*pb.FileInfo {
	var slice []*pb.FileInfo

	var i uint32
	for i = 0; i < number; i++ {
		slice = append(slice, NewFileInfo())
	}

	return slice
}

// NewRequest generates new request message
func NewRequest(number uint32) *pb.Request {
	req := &pb.Request{
		User:            NewCredentials(),
		NumberOfFiles:   number,
		FileInformation: NewFilesSlice(number),
	}

	return req
}

// NewChunk generates new chunk from the data
func NewChunk(bytes []byte) *pb.Chunk {
	chunk := &pb.Chunk{
		Content: bytes,
	}

	return chunk
}
