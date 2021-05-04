package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sheerley/dsim/codes"
	"github.com/Sheerley/dsim/convo"
	"github.com/Sheerley/dsim/pb"
	"github.com/Sheerley/dsim/plog"
	"github.com/Sheerley/dsim/transfer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	configLocation := os.Getenv("CONFIG")
	if len(configLocation) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			plog.Fatalf(codes.ConfError, "unable to dine home folder: %v", err)
		}
		configLocation = home + "/.config/dsim.d/config_client.json"
		plog.Messagef("config env variable not set, current config location: %v", configLocation)
	}

	createUser := flag.Bool("user-new", false, "set true if you want to create new user")
	deleteUser := flag.Bool("user-del", false, "set true if you want to delete existing user")
	modifyUser := flag.Bool("user-mod", false, "set true if you want to modify user")
	login := flag.String("uname", "", "username")
	passphrase := flag.String("pwd", "", "passphrase")
	newPassphrase := flag.String("pwd-new", "", "new passphrase")
	srcImg1 := flag.String("source-img1", "", "path to first source image file")
	srcImg2 := flag.String("source-img2", "", "path to second source image file")
	outFile := flag.String("o", "result", "path to output file, should not contain extension")

	logDescription := fmt.Sprintf("log level with possible values:\n - Verbose: %v\n - Debug: %v\n - Info: %v"+
		"\n - Waring: %v\n - Error: %v not recommended\n - Fatal: %v not recommended\n",
		plog.VERBOSE, plog.DEBUG, plog.INFO, plog.WARNING, plog.ERROR, plog.FATAL)
	logLevel := flag.Int("log-level", plog.WARNING, logDescription)

	flag.Parse()

	if len(*login) == 0 || len(*passphrase) == 0 {
		plog.Fatalf(codes.IncorrectArgs, "passphrase or username not provided")
	}

	if *createUser == false && *deleteUser == false && *modifyUser == false && (len(*srcImg1) == 0 || len(*srcImg2) == 0) {
		plog.Fatalf(codes.IncorrectArgs, "srcImg1 or srcImg2 path not provided")
	}

	plog.SetLogLevel(*logLevel)

	conf, err := convo.LoadConfiguration(configLocation)
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	address := fmt.Sprintf("%v:%v", conf.Address, conf.Port)

	plog.Messagef("dial server %v", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		plog.Fatalf(codes.ClientConnectionError, "cannot dial server: %v", err)
	}

	if *createUser == true {
		umClient := pb.NewUserServiceClient(conn)

		creds := transfer.NewCredentials(*login, *passphrase)
		req := &pb.ActionUserRequest{
			Credentials: creds,
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		res, err := umClient.CreateUser(ctx, req)

		if err != nil {
			st, ok := status.FromError(err)
			if ok && pb.Response_ReturnCode(st.Code()) == pb.Response_error {
				plog.Fatalf(codes.DbError, "%v", err)
			}
		}

		plog.Messagef("created user: %v", res.Response.ReturnCode)
	} else if *modifyUser == true {
		umClient := pb.NewUserServiceClient(conn)

		oldCreds := transfer.NewCredentials(*login, *passphrase)
		newCreds := transfer.NewCredentials(*login, *newPassphrase)
		req := &pb.ModifyUserRequest{
			OldCredentials: oldCreds,
			NewCredentials: newCreds,
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		res, err := umClient.ModifyUser(ctx, req)

		if err != nil {
			st, ok := status.FromError(err)
			if ok && pb.Response_ReturnCode(st.Code()) == pb.Response_error {
				plog.Fatalf(codes.DbError, "%v", err)
			}
		}

		plog.Messagef("modified user: %v", res.Response.ReturnCode)
	} else if *deleteUser == true {
		umClient := pb.NewUserServiceClient(conn)

		creds := transfer.NewCredentials(*login, *passphrase)
		req := &pb.ActionUserRequest{
			Credentials: creds,
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		res, err := umClient.DeleteUser(ctx, req)

		if err != nil {
			st, ok := status.FromError(err)
			if ok && pb.Response_ReturnCode(st.Code()) == pb.Response_error {
				plog.Fatalf(codes.DbError, "%v", err)
			}
		}

		plog.Messagef("deleted user: %v", res.Response.ReturnCode)
	} else {
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

		if len(*srcImg1) > 0 {
			filenames = append(filenames, *srcImg1)
		}
		if len(*srcImg2) > 0 {
			filenames = append(filenames, *srcImg2)
		}

		for i := 0; i < len(filenames); i++ {
			f, err := os.Open(filenames[i])
			if err != nil {
				plog.Fatalf(codes.FileError, "%v\n", err)
			}

			files = append(files, f)

			var buffer []byte
			chunk := make([]byte, 1024)
			for {
				n, err := f.Read(chunk)
				if err == io.EOF {
					break
				}
				if err != nil {
					plog.Fatalf(codes.FileError, "checksum calculation file reading: %v", err)
				}
				buffer = append(buffer, chunk[:n]...)
			}

			s256 := sha256.New()
			checksum := s256.Sum(buffer)

			fo := &pb.FileInfo{
				FileExtension: filepath.Ext(filenames[i]),
				Checksum:      checksum,
			}

			fileinfo = append(fileinfo, fo)

			f.Seek(0, io.SeekStart)

			defer files[len(files)-1].Close()
		}

		job := &pb.Job{
			User:            transfer.NewCredentials(*login, *passphrase),
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

		res, err := stream.Recv()
		if err != nil {
			plog.Fatalf(codes.ServerError, "unable to process request: \n- %v \n- %v", err, stream.RecvMsg(nil))
		}

		skipped := res.GetResponse().GetContext()
		if len(skipped) > 0 {
			plog.Verbosef("skipped %v files: %v", len(skipped), skipped)
		}

		for i := 0; i < len(files); i++ {
			cont := false
			for _, skip := range skipped {
				if skip == int32(i) {
					plog.Verbosef("skipping %v", i)
					cont = true
					break
				}
			}

			if cont {
				continue
			}

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

		res, err = stream.Recv()
		for {
			if err != nil {
				plog.Fatalf(codes.ServerError, "cannot recieve response: %v", err)
			}

			if res != nil {
				plog.Debugf("response: %v", res.GetResponse().GetReturnMessage())
				break
			} else {
				res, err = stream.Recv()
			}
		}

		var recievedFile []byte

		fileData := bytes.Buffer{}
		fileSize := 0

		if res.GetResponse().GetReturnCode() != pb.Response_ok {
			plog.Fatalf(codes.ServerError, "failed to finish the job: \n- %v", res.GetResponse().GetReturnMessage())
		}

		res, err = stream.Recv()
		if err != nil {
			plog.Errorf("failed to recieve file info: %v", err)
		}

		extension := res.GetFileInfo().GetFileExtension()
		plog.Debugf("file extension recieved: %v", extension)

		for {
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

		err = ioutil.WriteFile(*outFile+extension, recievedFile, 0644)
		if err != nil {
			plog.Fatalf(codes.FileError, "unable to write file: %v", err)
		}

		plog.Messagef("Done! Result written to %v", *outFile+extension)
	}

}
