package main

import (
	"context"
	"fmt"

	"github.com/Sheerley/pluggabl/pkg/transfer"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

	umClient := pb.NewUserServiceClient(conn)

	creds := transfer.NewDummyCredentials()
	req := &pb.ActionUserRequest{
		Credentials: creds,
	}

	res, err := umClient.CreateUser(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && pb.Response_ReturnCode(st.Code()) == pb.Response_error {
			plog.Fatalf(codes.DbError, "%v %v", err)
		}
	}

	plog.Messagef("created user: %v", res.Response.ReturnCode)
}
