GOCV_VERSION=v0.25.0 

.PHONY: build preinstall install clean proto test 

all: preinstall clean proto build

preinstall:
	go mod tidy
	cd $(shell go env GOPATH)/pkg/mod/gocv.io/x/gocv\@$(GOCV_VERSION) && $(MAKE) install

clean:
	rm -f ./pkg/pb/*.pb.go
	rm -f ./build/*

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

build:
	go build -o ./build/pluggabl-primary ./cmd/server/primary 
	go build -o ./build/pluggabl-secondary ./cmd/server/secondary 
	go build -o ./build/pluggabl-exec ./cmd/exec 
	go build -o ./build/pluggabl-client ./cmd/client

install:
	sudo cp ./build/pluggabl-* /tmp/
	[ ! -d ~/.config/pluggabl.d/ ] && mkdir ~/.config/pluggabl.d/ || echo ok
	cp ./config/*.json ~/.config/pluggabl.d/

test:
	go test -cover -race ./...
