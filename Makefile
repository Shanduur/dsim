GOCV_VERSION=v0.24.0 

all: install clean proto gtest

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

clean:
	rm -f pkg/pb/*.pb.go
	rm -f */**/*.tmp.*

server:
	go run ./cmd/server

client:
	go run ./cmd/client

test:
	go test -cover -race ./...

ptest: proto test

install:
	cd $(shell go env GOPATH)/pkg/mod/gocv.io/x/gocv\@$(GOCV_VERSION) && $(MAKE) install