package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sheerley/pluggabl/pkg/transfer"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
	"google.golang.org/grpc"
)

func main() {
	plog.SetLogLevel(plog.VERBOSE)

	conf, err := convo.LoadConfiguration("config/config_client.json")
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	address := fmt.Sprintf("%v:%v", conf.ManagerAddress, conf.ManagerPort)

	plog.Messagef("dial server %v", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		plog.Fatalf(codes.ClientConnectionError, "cannot dial server: %v", err)
	}

	jobClient := pb.NewJobServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := jobClient.SubmitJob(ctx)
	if err != nil {
		plog.Fatalf(codes.ClientConnectionError, "unable to create stream: %v", err)
	}

	var filenames []string
	var files []*os.File
	var fileinfo []*pb.FileInfo

	filenames = append(filenames, "examples/images/box.png")
	filenames = append(filenames, "examples/images/box_in_scene.png")

	for i := 0; i < len(filenames); i++ {
		f, err := os.Open(filenames[i])
		if err != nil {
			plog.Fatalf(codes.FileError, "%v\n", err)
		}

		files = append(files, f)

		fo := &pb.FileInfo{
			FileExtension: filepath.Ext(filenames[i]),
		}

		fileinfo = append(fileinfo, fo)

		defer files[len(files)-1].Close()
	}

	job := &pb.Job{
		User:            transfer.NewDummyCredentials(),
		NumberOfFiles:   uint32(len(files)),
		FileInformation: fileinfo,
	}

	req := &pb.JobRequest{
		Data: &pb.JobRequest_Job{
			Job: job,
		},
	}

	err = stream.Send(req)
	if err != nil {
		plog.Fatalf(codes.ManagerConnectionError, "unable to process request: \n- %v \n- %v", err, stream.RecvMsg(nil))
	}

	for i := 0; i < len(files); i++ {
		reader := bufio.NewReader(files[i])

		buffer := make([]byte, 1024)

		plog.Debugf("file %v", i)

		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				plog.Fatalf(codes.FileError, "error reading file: %v", err)
			}

			plog.Verbosef("sending data to server for file %v", i)

			req := &pb.JobRequest{
				Data: &pb.JobRequest_ChunkData{
					ChunkData: &pb.Chunk{
						FileNumber: int32(i),
						Content:    buffer[:n],
					},
				},
			}

			err = stream.Send(req)
			if err != nil {
				plog.Fatalf(codes.ServerError, "cannot send chunk to server: \n- %v \n- %v", err, stream.RecvMsg(nil))
			}
		}
	}

	plog.Debugf("finished sending")

	stream.CloseSend()
	if err != nil {
		plog.Errorf("unable to close send %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		plog.Fatalf(codes.ServerError, "cannot recieve response: %v", err)
	}

	var recievedFile []byte

	fileData := bytes.Buffer{}
	fileSize := 0

	if res.GetResponse().GetReturnCode() != pb.Response_ok {
		plog.Fatalf(codes.ServerError, "failed to finish the job: %v", err)
	}

	res, err = stream.Recv()
	if err != nil {
		plog.Errorf("failed to recieve file info: %v", err)
	}

	extension := res.GetFileInfo().GetFileExtension()
	plog.Debugf("file extension recieved: %v", extension)

	for {
		plog.Verbosef("waiting to recieve more data for result file")

		res, err = stream.Recv()
		if err == io.EOF {
			recievedFile = fileData.Bytes()

			plog.Debugf("size of file: %v", fileSize)

			plog.Messagef("recieving finished")

			break
		}

		if err != nil {
			plog.Errorf("cannot recieve chunk data: %v", err)

			return
		}

		chunk := res.GetChunkData().GetContent()
		size := len(chunk)

		fileSize += size

		_, err = fileData.Write(chunk)
		if err != nil {
			plog.Errorf("cannot write chunk data: %v", err)

			return
		}
	}

	if len(recievedFile) == 0 {
		plog.Fatalf(codes.ServerError, "failed to recieve file")
	}

	err = ioutil.WriteFile("recieved.tmp"+extension, recievedFile, 0644)
	if err != nil {
		plog.Fatalf(codes.FileError, "unable to write file: %v", err)
	}

	plog.Messagef("Done! Result written to %v", "recieved.tmp"+extension)
}
