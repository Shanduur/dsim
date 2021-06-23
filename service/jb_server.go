package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/Sheerley/dsip/convo"
	"github.com/Sheerley/dsip/db"
	"github.com/Sheerley/dsip/pb"
	"github.com/Sheerley/dsip/plog"
	"github.com/google/uuid"
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

	cfg := convo.SavedConfig

	err := db.UpdateJobStatus(cfg, +1)
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

		if ext == "unknown" {
			extension = ".unn"
		} else {
			extension = ext
		}

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

	outname := uuid.New().String() + extension
	filenames = append(filenames, outname)

	defer purgeFiles(filenames)

	var argv []string
	for _, f := range filenames {
		arg := fmt.Sprintf("-img=%v", f)

		argv = append(argv, arg)
	}

	argv = append(argv, "-out="+outname)

	job := exec.CommandContext(ctx, cfg.JobBinaryName, argv...)

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

	var skipped []int32
	id, err := db.UploadResult(ctx, fileSlice, skipped, fileInfo, fileIDs)
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

func purgeFiles(fnames []string) {
	for _, name := range fnames {
		plog.Debugf("removing file %v", name)
		err := os.Remove(name)
		if err != nil {
			plog.Warningf("encountered problem while removing file %v: %v", name, err)
		}
	}
}
