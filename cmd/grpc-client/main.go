package main

import (
	"context"

	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/transfer"
)

func main() {
	var client transfer.Client

	grpcClient, err := transfer.NewClientGRPC(transfer.ClientGRPCConfig{
		Address:         "127.0.0.1:5070",
		RootCertificate: "",
		ChunkSize:       8192,
	})

	if err != nil {
		plog.Fatalf(1, "failed to create client: %v", err)
	}

	client = &grpcClient

	err = client.UploadFile(context.Background(), "./cmd/grpc-client/file.txt")

	if err != nil {
		plog.Fatalf(1, "Upload failed: %v", err)
	}

	defer client.Close()
}
