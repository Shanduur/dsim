GOCV_VERSION=v0.25.0 

all: install clean proto gtest build

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

compile:
	go build -o ./build/primary ./cmd/server/primary 
	go build -o ./build/secondary ./cmd/server/secondary 
	go build -o ./build/exec ./cmd/exec 

clean:
	rm -f pkg/pb/*.pb.go
	rm -f */**/*.tmp.*

server-p:
	go run ./cmd/server/primary

server-s:
	go run ./cmd/server/secondary

client-u:
	go run ./cmd/client/um

client-t:
	go run ./cmd/client/tr

test:
	go test -cover -race ./...

ptest: proto test

install:
	go mod tidy
	cd $(shell go env GOPATH)/pkg/mod/gocv.io/x/gocv\@$(GOCV_VERSION) && $(MAKE) install