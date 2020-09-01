package transfer

import (
	"crypto/sha512"

	"github.com/Sheerley/pluggabl/pkg/pb"
)

// NewDummyCredentials generates Credentials struct
func NewDummyCredentials() *pb.Credentials {
	s512 := sha512.New()
	bv := []byte("placeholder")

	userCreds := &pb.Credentials{
		UserId:  "placeholder",
		UserKey: s512.Sum(bv),
	}

	return userCreds
}

// NewDummyFileInfo generates File information
func NewDummyFileInfo() *pb.FileInfo {
	fileInfo := &pb.FileInfo{
		FileExtension: ".txt",
		SizeType:      pb.FileInfo_kilobytes,
		Size:          2048,
	}

	return fileInfo
}

// NewDummyFilesSlice generates slice containing informations about transfered files
func NewDummyFilesSlice(number uint32) []*pb.FileInfo {
	var slice []*pb.FileInfo

	var i uint32
	for i = 0; i < number; i++ {
		slice = append(slice, NewDummyFileInfo())
	}

	return slice
}

// NewDummyJob generates new request message
func NewDummyJob(number uint32) *pb.Job {
	req := &pb.Job{
		User:            NewDummyCredentials(),
		NumberOfFiles:   number,
		FileInformation: NewDummyFilesSlice(number),
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
