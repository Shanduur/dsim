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

	plog.SetLogLevel(plog.VERBOSE)

	conf, err := convo.LoadConfiguration("config/config_primary.json")
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	umServ := service.NewUserManagementServer()
	trServ := service.NewTransportServer()
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, umServ)
	pb.RegisterJobServiceServer(grpcServer, trServ)

	address := fmt.Sprintf("0.0.0.0:%v", conf.PrimaryNodePort)

	plog.Messagef("net.Listen tcp on %v", address)
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
