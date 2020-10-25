package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/google/uuid"

	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/db"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
)

// InternalJobServer is implementation of gRPC server
type InternalJobServer struct {
	mut sync.Mutex
}

// NewInternalJobServer returns instance of InternalJobServer
func NewInternalJobServer() *InternalJobServer {
	return &InternalJobServer{}
}

// SubmitJob function is
func (srv *InternalJobServer) SubmitJob(ctx context.Context, req *pb.InternalJobRequest) (*pb.InternalJobResponse, error) {
	var rsp *pb.InternalJobResponse

	cfg, err := convo.LoadConfiguration("default")
	if err != nil {
		err = fmt.Errorf("unable to retrieve file: %v", err)

		rsp = &pb.InternalJobResponse{
			Response: &pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			},
		}

		return rsp, err
	}

	db.UpdateJobStatus(cfg, +1)
	defer db.UpdateJobStatus(cfg, -1)
	defer plog.ContextStatus(ctx)

	fileIDs := req.GetJob().GetFileId()
	var filenames []string

	for i := 0; i < len(fileIDs); i++ {
		tmp, extension, err := db.GetFile(ctx, fileIDs[i])
		if err != nil {
			err = fmt.Errorf("unable to retrieve file: %v", err)

			rsp = &pb.InternalJobResponse{
				Response: &pb.Response{
					ReturnMessage: err.Error(),
					ReturnCode:    pb.Response_error,
				},
			}

			return rsp, err
		}

		name := uuid.New().String()
		name = name + ".tmp" + extension

		err = ioutil.WriteFile(name, tmp, 0644)
		if err != nil {
			err = fmt.Errorf("unable to write %v file: %v", name, err)

			rsp = &pb.InternalJobResponse{
				Response: &pb.Response{
					ReturnMessage: err.Error(),
					ReturnCode:    pb.Response_error,
				},
			}

			return rsp, err
		}

		filenames = append(filenames, name)
	}

	rsp = &pb.InternalJobResponse{
		Response: &pb.Response{
			ReturnCode: pb.Response_ok,
		},
	}

	return rsp, nil
}
