package main

import (
	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/transfer"
)

func main() {
	var (
		port        = 5070
		key         = ""
		certificate = ""
		server      transfer.Server
	)

	grpcServer, err := transfer.NewServerGRPC(transfer.ServerGRPCConfig{
		Port:        port,
		Certificate: certificate,
		Key:         key,
	})

	if err != nil {
		plog.Fatalf(1, "failed to create server: %v", err)
	}

	server = &grpcServer

	err = server.Listen()

	defer server.Close()
}
