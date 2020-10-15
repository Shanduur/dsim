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

			jrsp = &pb.JobResponse{
				Id:       append(id, codes.UnknownID),
				Response: rsp,
			}

			return
		}

		plog.Errorf("%v", err)

		jrsp = &pb.JobResponse{
			Id:       append(id, codes.UnknownID),
			Response: rsp,
		}

		return
	}

	fileInfoTabSize := uint32(len(fileInfo))

	if fileInfoTabSize != numberOfFiles {
		plog.Errorf("%v", err)

		return
	}

	// create temporary storage for blob
	tempStorage := make([][]byte, numberOfFiles)

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
			plog.Errorf("cannot recieve chunk data: %v", err)

			return
		}

		chunk := req.GetChunkData().GetContent()
		fileNum := req.GetChunkData().GetFileNumber()
		size := len(chunk)

		if fileNum != currentFile {
			// copy data to temp storage
			tempStorage[currentFile] = make([]byte, len(fileData.Bytes()))
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
			plog.Errorf("cannot write chunk data: %v", err)

			return
		}
	}

	tempID, err := db.UploadFiles(ctx, tempStorage, fileInfo, user)

	for i := 0; i < len(tempID); i++ {
		id = append(id, int64(tempID[i]))
	}

	if err != nil {
		msg := fmt.Sprintf("cannot upload file to database: %v", err)
		rsp := &pb.Response{
			ReturnMessage: msg,
			ReturnCode:    pb.Response_error,
		}

		jrsp = &pb.JobResponse{
			Id:       id,
			Response: rsp,
		}

		plog.Errorf("%v", err)

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
