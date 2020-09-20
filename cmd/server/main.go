package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

	err := convo.LoadConfiguration("default")
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	umServ := service.NewUserManagementServer()
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, umServ)

	address := fmt.Sprintf("0.0.0.0:%v", convo.Configuration.PrimaryNodePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		plog.Fatalf(codes.ServerError, "error while creating listener: %v", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(listener)
		if err != nil {
			plog.Fatalf(codes.ServerError, "error while registering listener %v", err)
		}
	}()

	wg.Wait()
}
