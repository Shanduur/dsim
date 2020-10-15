package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

	var filenames []string
	var files []*os.File
	var fileinfo []*pb.FileInfo

	filenames = append(filenames, "box.png")
	filenames = append(filenames, "box_in_scene.png")

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

		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				plog.Fatalf(codes.FileError, "error reading file: %v", err)
			}

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

	res, err := stream.CloseAndRecv()
	if err != nil {
		plog.Fatalf(codes.ServerError, "cannot recieve response: %v", err)
	}

	plog.Messagef("job request succesfully sent: file ID = %v", res.GetId())
}
