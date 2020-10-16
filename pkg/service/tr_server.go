package service

import (
	"bytes"
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
func (srv *TransportServer) SubmitJob(stream pb.JobService_SubmitJobServer) (err error) {
	ctx := stream.Context()
	defer plog.ContextStatus(ctx)

	var jrsp *pb.JobResponse

	var id []int64

	req, err := stream.Recv()
	if err != nil {
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
			plog.Errorf("%v", err)

			return
		}

		err = fmt.Errorf("unable to process request")
		plog.Errorf("%v", err)

		return
	}

	fileInfoTabSize := uint32(len(fileInfo))

	if fileInfoTabSize != numberOfFiles {
		err = fmt.Errorf("unable to process request - data mismatch")

		plog.Errorf("%v", err)

		return
	}

	// create temporary storage for blob
	tempStorage := make([][]byte, numberOfFiles)

	fileData := bytes.Buffer{}
	fileSize := 0

	currentFile := int32(0)

	for {
		plog.Verbosef("waiting to recieve more data from file number %v", currentFile)

		req, err = stream.Recv()
		if err == io.EOF {
			tempStorage[currentFile] = make([]byte, len(fileData.Bytes()))
			tempStorage[currentFile] = fileData.Bytes()

			plog.Debugf("size of file %v: %v", currentFile, fileSize)

			plog.Messagef("recieving finished")

			break
		}

		if err != nil {
			plog.Errorf("cannot recieve chunk data: %v", err)

			return
		}

		chunk := req.GetChunkData().GetContent()
		fileNum := req.GetChunkData().GetFileNumber()
		size := len(chunk)

		if fileNum != currentFile {
			plog.Debugf("changing file from %v to %v", currentFile, fileNum)
			// copy data to temp storage
			tempStorage[currentFile] = make([]byte, len(fileData.Bytes()))
			tempStorage[currentFile] = fileData.Bytes()

			// empty the data
			fileData.Reset()

			plog.Debugf("size of file %v: %v", currentFile, fileSize)
			// reset file size
			fileSize = 0

			// change file num
			currentFile = fileNum
		}

		fileSize += size

		_, err = fileData.Write(chunk)
		if err != nil {
			plog.Errorf("cannot write chunk data: %v", err)

			return
		}
	}

	tempID, err := db.UploadFiles(ctx, tempStorage, fileInfo, user)

	for i := 0; i < len(tempID); i++ {
		id = append(id, int64(tempID[i]))
	}

	if err != nil {
		plog.Errorf("cannot upload file to database: %v", err)

		return
	}

	msg := fmt.Sprintf("succesfully uploaded images with id: %v", id)

	rsp := &pb.Response{
		ReturnMessage: msg,
		ReturnCode:    pb.Response_ok,
	}

	jrsp = &pb.JobResponse{
		Id:       id,
		Response: rsp,
	}

	err = stream.SendAndClose(jrsp)
	if err != nil {
		plog.Errorf("%v", err)

		return
	}

	plural := ""
	if len(id) > 1 {
		plural = "s"
	}
	plog.Messagef("succesfully recieved %v blob%v with id%v: %v", len(id), plural, plural, id)

	return
}
