package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/Sheerley/pluggabl/internal/codes"
	"github.com/Sheerley/pluggabl/internal/convo"
	"github.com/Sheerley/pluggabl/internal/fuse"
	"github.com/Sheerley/pluggabl/pkg/db"
	"github.com/Sheerley/pluggabl/pkg/pb"
	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

	plog.SetLogLevel(plog.VERBOSE)

	conf, err := convo.LoadConfiguration("config/config_secondary.json")
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	ijServ := service.NewInternalJobServer()
	grpcServer := grpc.NewServer()

	pb.RegisterInternalJobServiceServer(grpcServer, ijServ)

	address := fmt.Sprintf("0.0.0.0:%v", conf.SecondaryNodePort)

	plog.Messagef("net.Listen tcp on %v", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		plog.Fatalf(codes.ServerError, "error while creating listener: %v", err)
	}

	err = db.RegisterNode(conf)
	if err != nil {
		plog.Fatalf(codes.DbError, "unable to register node: %v", err)
	}
	defer db.UnregisterNode(conf)

	wg.Add(1)
	go fuse.Watchdog(&wg)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(listener)
		if err != nil {
			plog.Fatalf(codes.ServerError, "error while registering listener: %v", err)
		}
	}()

	wg.Wait()
}
