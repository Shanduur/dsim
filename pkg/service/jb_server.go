package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/Sheerley/pluggabl/pkg/transfer"

	"github.com/google/uuid"

	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/db"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
)

// InternalJobServer is implementation of gRPC server
type InternalJobServer struct{}

// NewInternalJobServer returns instance of InternalJobServer
func NewInternalJobServer() *InternalJobServer {
	return &InternalJobServer{}
}

// SubmitJob function is
func (srv *InternalJobServer) SubmitJob(ctx context.Context, req *pb.InternalJobRequest) (*pb.InternalJobResponse, error) {
	var rsp *pb.InternalJobResponse

	cfg, err := convo.LoadConfiguration("config/config_secondary.json")
	if err != nil {
		err = fmt.Errorf("unable to retrieve file: %v", err)

		rsp = &pb.InternalJobResponse{
			Response: &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			},
		}

		plog.Errorf("%v", err)

		return rsp, err
	}

	defer db.UpdateJobStatus(cfg, -1)
	defer plog.ContextStatus(ctx)

	fileIDs := req.GetJob().GetFileId()
	var filenames []string

	var extension string

	for i := 0; i < len(fileIDs); i++ {
		tmp, ext, err := db.GetFile(ctx, fileIDs[i])
		if err != nil {
			err = fmt.Errorf("unable to retrieve file: %v", err)

			rsp = &pb.InternalJobResponse{
				Response: &pb.Response{
					ReturnMessage: err.Error(),
					ReturnCode:    pb.Response_error,
				},
			}

			plog.Errorf("%v", err)

			return rsp, err
		}

		extension = ext

		name := uuid.New().String()
		name = name + ".tmp" + ext

		err = ioutil.WriteFile(name, tmp, 0644)
		if err != nil {
			err = fmt.Errorf("unable to write %v file: %v", name, err)

			rsp = &pb.InternalJobResponse{
				Response: &pb.Response{
					ReturnMessage: err.Error(),
					ReturnCode:    pb.Response_error,
				},
			}

			plog.Errorf("%v", err)

			return rsp, err
		}

		filenames = append(filenames, name)
	}

	args := ""

	for i := 0; i < len(filenames); i++ {
		args += filenames[i] + " "
	}

	outname := uuid.New().String() + ".tmp" + extension
	job := exec.Command(cfg.JobBinaryName, args, "-o "+outname)

	err = job.Run()
	if err != nil {
		err = fmt.Errorf("unable to process the job: %v", err)

		rsp = &pb.InternalJobResponse{
			Response: &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			},
		}

		plog.Errorf("%v", err)

		return rsp, err
	}

	outfile, err := ioutil.ReadFile(outname)
	if err != nil {
		err = fmt.Errorf("unable to read the output file: %v", err)

		rsp = &pb.InternalJobResponse{
			Response: &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			},
		}

		plog.Errorf("%v", err)

		return rsp, err
	}

	var fileSlice [][]byte
	fileSlice = append(fileSlice, outfile)

	var fileInfo []*pb.FileInfo
	fileInfo = append(fileInfo, &pb.FileInfo{
		FileExtension: extension,
	})

	creds := transfer.NewAdminCredentials()

	id, err := db.UploadFiles(ctx, fileSlice, fileInfo, creds)
	if err != nil {
		err = fmt.Errorf("unable to upload the output file: %v", err)

		rsp = &pb.InternalJobResponse{
			Response: &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			},
		}

		plog.Errorf("%v", err)

		return rsp, err
	}

	rsp = &pb.InternalJobResponse{
		FileId: id,
		Response: &pb.Response{
			ReturnCode: pb.Response_ok,
		},
	}

	return rsp, nil
}
