package service

import (
	"bytes"
	"fmt"
	"io"

	"github.com/Sheerley/pluggabl/internal/convo"
	"google.golang.org/grpc"

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
			e := copy(tempStorage[currentFile], fileData.Bytes())

			plog.Verbosef("elements copied: %v", e)

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
			e := copy(tempStorage[currentFile], fileData.Bytes())

			plog.Verbosef("elements copied: %v", e)

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

	plural := ""
	if len(id) > 1 {
		plural = "s"
	}
	plog.Messagef("succesfully recieved %v blob%v with id%v: %v", len(id), plural, plural, id)

	conf, err := convo.LoadConfiguration("config/config_primary.json")
	if err != nil {
		msg := fmt.Sprintf("cannot read config: %v", err)

		jrsp = &pb.JobResponse{
			Data: &pb.JobResponse_Response{
				Response: &pb.Response{
					ReturnCode:    pb.Response_error,
					ReturnMessage: msg,
				},
			},
		}

		err2 := stream.Send(jrsp)

		plog.Errorf("cannot read config: \n- %v\n- %v", err, err2)
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.SecondaryNodeAddress, conf.SecondaryNodePort)

	plog.Messagef("dial secondary node %v", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		msg := fmt.Sprintf("cannot dial secondary node: %v", err)

		jrsp = &pb.JobResponse{
			Data: &pb.JobResponse_Response{
				Response: &pb.Response{
					ReturnCode:    pb.Response_error,
					ReturnMessage: msg,
				},
			},
		}

		err2 := stream.Send(jrsp)

		plog.Errorf("cannot dial secondary node: \n- %v\n- %v", err, err2)
		return err
	}

	ijClient := pb.NewInternalJobServiceClient(conn)

	ireq := &pb.InternalJobRequest{
		Job: &pb.InternalJob{
			FileId: id,
		},
	}

	res, err := ijClient.SubmitJob(ctx, ireq)
	if err != nil {
		msg := fmt.Sprintf("unable to finish the job: %v", err)

		jrsp = &pb.JobResponse{
			Data: &pb.JobResponse_Response{
				Response: &pb.Response{
					ReturnCode:    pb.Response_error,
					ReturnMessage: msg,
				},
			},
		}

		err2 := stream.Send(jrsp)

		plog.Errorf("error submitting job: \n- %v\n- %v", err, err2)
		return err
	}

	fileID := res.GetFileId()

	if len(fileID) > 1 {
		err = fmt.Errorf("unexpected number of return files: %v", len(fileID))
		plog.Errorf(err.Error())
		return err
	}
	resultFile, extension, err := db.GetFile(ctx, fileID[0])
	if err != nil {
		plog.Errorf("unable to download file from database: %v", err)
		return err
	}

	jrsp = &pb.JobResponse{
		Data: &pb.JobResponse_FileInfo{
			FileInfo: &pb.FileInfo{
				FileExtension: extension,
			},
		},
	}

	err = stream.Send(jrsp)
	if err != nil {
		plog.Errorf("cannot send file info to client: \n- %v \n- %v", err, stream.RecvMsg(nil))
		return err
	}

	reader := bytes.NewReader(resultFile)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			msg := fmt.Sprintf("error reading file: %v", err)

			jrsp = &pb.JobResponse{
				Data: &pb.JobResponse_Response{
					Response: &pb.Response{
						ReturnCode:    pb.Response_error,
						ReturnMessage: msg,
					},
				},
			}

			err2 := stream.Send(jrsp)

			plog.Errorf("error submitting job: \n- %v\n- %v", err, err2)
			return err
		}

		plog.Verbosef("sending result data to client")

		jrsp = &pb.JobResponse{
			Data: &pb.JobResponse_ChunkData{
				ChunkData: &pb.Chunk{
					Content: buffer[:n],
				},
			},
		}

		err = stream.Send(jrsp)
		if err != nil {
			plog.Errorf("cannot send chunk to client: \n- %v \n- %v", err, stream.RecvMsg(nil))
			return err
		}
	}

	return
}
