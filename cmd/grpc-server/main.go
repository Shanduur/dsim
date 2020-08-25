package main

import (
	"net"

	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/transfer"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		plog.Fatalf(1, "failed to listen: %v\n", err)
	}

	s := transfer.Server{}

	grpcServer := grpc.NewServer()

	transfer.RegisterTransferServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		plog.Fatalf(1, "failed to serve: %v\n", err)
	}
}
