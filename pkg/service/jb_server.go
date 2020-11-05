package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

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

	err = db.UpdateJobStatus(cfg, +1)
	if err != nil {
		err = fmt.Errorf("unable to update job status: %v", err)

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
		tmp, name, ext, err := db.GetFile(ctx, fileIDs[i])
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

		name = name + ext

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

	query := "-query=" + filenames[0]
	train := "-train=" + filenames[1]

	outname := uuid.New().String() + extension
	job := exec.Command(cfg.JobBinaryName, query, train, "-out="+outname)

	c := make(chan bool)

	plog.Messagef("job started")

	go func(c <-chan bool) {
		for {
			select {
			case <-c:
				plog.Messagef("job done")
				return
			default:
				plog.Verbosef("working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(c)

	err = job.Run()
	c <- true
	close(c)
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
