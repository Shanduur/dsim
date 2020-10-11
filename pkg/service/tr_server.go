package service

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/pkg/db"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
)

// TransportServer struct is implementation of the gRPC server
type TransportServer struct {
}

// NewTransportServer function initializes new server for the gRPC
func NewTransportServer() *TransportServer {
	return &TransportServer{}
}

// SubmitJob is handler function for requesting Job
func (srv *TransportServer) SubmitJob(ctx context.Context, stream pb.JobService_SubmitJobServer) (rsp *pb.Response, err error) {
	defer plog.ContextStatus(ctx)

	req, err := stream.Recv()
	if err != nil {
		rsp = &pb.Response{
			ReturnMessage: err.Error(),
			ReturnCode:    pb.Response_error,
		}

		err = fmt.Errorf("cannot recieve file info: %d", err)

		plog.Errorf("%v", err)

		return
	}

	job := req.GetJob()
	user := job.GetUser()
	numberOfFiles := job.GetNumberOfFiles()
	fileInfo := job.GetFileInformation()

	err = db.Auth(ctx, user)
	if err != nil {
		if err.Error() == (&codes.NotAuthenticated{}).Error() {
			rsp = &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			return
		}

		rsp = &pb.Response{
			ReturnMessage: "unable to process request",
			ReturnCode:    pb.Response_unknown,
		}

		return
	}

	fileInfoTabSize := uint32(len(fileInfo))

	if fileInfoTabSize != numberOfFiles {
		rsp = &pb.Response{
			ReturnMessage: "unable to process request - data mismatch",
			ReturnCode:    pb.Response_error,
		}

		return
	}

	var size uint64
	size = 0

	// get max size
	for _, f := range fileInfo {
		if size < f.GetSize() {
			size = f.GetSize()
		}
	}

	// create temporary storage for blob
	tempStorage := make([][]byte, numberOfFiles)
	for i := uint32(0); i < numberOfFiles; i++ {
		tempStorage[i] = make([]byte, size)
	}

	fileData := bytes.Buffer{}
	fileSize := 0

	currentFile := int32(0)

	for {
		plog.Messagef("waiting to recieve more data")

		req, err = stream.Recv()
		if err == io.EOF {
			plog.Messagef("recieving finished")
			break
		}

		if err != nil {
			msg := fmt.Sprintf("cannot recieve chunk data: %v", err)
			rsp = &pb.Response{
				ReturnMessage: msg,
				ReturnCode:    pb.Response_unknown,
			}

			return
		}

		chunk := req.GetChunkData().GetContent()
		fileNum := req.GetChunkData().GetFileNumber()
		size := len(chunk)

		if fileNum != currentFile {
			// copy data to temp storage
			tempStorage[currentFile] = fileData.Bytes()

			// empty the data
			fileData.Reset()

			// change file num
			currentFile = fileNum

			// reset file size
			fileSize = 0
		}

		fileSize += size

		_, err = fileData.Write(chunk)
		if err != nil {
			msg := fmt.Sprintf("cannot write chunk data: %v", err)
			rsp = &pb.Response{
				ReturnMessage: msg,
				ReturnCode:    pb.Response_error,
			}

			return
		}
	}

	return
}
