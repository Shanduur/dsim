package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/convo"
	"github.com/Sheerley/pluggabl/fuse"
	"github.com/Sheerley/pluggabl/pb"
	"github.com/Sheerley/pluggabl/plog"
	"github.com/Sheerley/pluggabl/service"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

	plog.Warningf("PID: %v", os.Getpid())

	configLocation := os.Getenv("CONFIG")
	if len(configLocation) == 0 {
		configLocation = "/etc/pluggabl/config.json"
		plog.Warningf("config env variable not set, current config location: %v", configLocation)
	}

	logDescription := fmt.Sprintf("log level with possible values:\n - Verbose: %v\n - Debug: %v\n - Info: %v"+
		"\n - Waring: %v\n - Error: %v not recommended\n - Fatal: %v not recommended",
		plog.VERBOSE, plog.DEBUG, plog.INFO, plog.WARNING, plog.ERROR, plog.FATAL)
	logLevel := flag.Int("log-level", plog.WARNING, logDescription)

	flag.Parse()

	plog.SetLogLevel(*logLevel)

	conf, err := convo.LoadConfiguration(configLocation)
	if err != nil {
		plog.Fatalf(codes.ConfError, "error while loading configuration: %v", err)
	}

	plog.Splash(conf.Tell())
	plog.Verbose(conf)
	plog.Verbose(os.Getenv("PG_DATABASE"))

	umServ := service.NewUserManagementServer()
	trServ := service.NewTransportServer()
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, umServ)
	pb.RegisterJobServiceServer(grpcServer, trServ)

	address := fmt.Sprintf("0.0.0.0:%v", conf.Port)

	plog.Messagef("net.Listen tcp on %v", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		plog.Fatalf(codes.ServerError, "error while creating listener: %v", err)
	}

	wg.Add(1)
	go fuse.Watchdog(&wg)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(listener)
		if err != nil {
			plog.Fatalf(codes.ServerError, "error while registering listener %v", err)
		}
	}()

	wg.Wait()
}
