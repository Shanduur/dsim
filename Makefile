GOCV_VERSION=v0.25.0 

.PHONY: build prebuild install clean proto test docker

all: prebuild clean proto build

prebuild:
	go mod download
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
	go build -o ./build/pluggabl-primary 	./cmd/server/primary 
	go build -o ./build/pluggabl-secondary 	./cmd/server/secondary 
	go build -o ./build/pluggabl-exec		./cmd/exec 
	go build -o ./build/pluggabl-client 	./cmd/client

install:
	[ ! -d /opt/pluggabl/ ] && sudo mkdir /opt/pluggabl/ || echo ok

	sudo cp ./build/pluggabl-primary	/opt/pluggabl/primary.run
	sudo cp ./build/pluggabl-secondary	/opt/pluggabl/secondary.run
	sudo cp ./build/pluggabl-exec		/opt/pluggabl/exec.run
	sudo cp ./build/pluggabl-client		/opt/pluggabl/client.run
	sudo cp ./scripts/pluggabl.sh		/opt/pluggabl/pluggabl.sh

	sudo ln -sf /opt/pluggabl/pluggabl.sh		/usr/bin/pluggabl
	
	sudo chmod +x /opt/pluggabl/*

	[ ! -d /etc/pluggabl/ ] && sudo mkdir /etc/pluggabl/ || echo ok
	sudo cp ./config/*.json /etc/pluggabl/

test:
	go test -cover -race ./...
	
docker: build
	cp ./build/pluggabl-primary 	./docker/primary/files/pluggabl/server.run
	cp ./build/pluggabl-secondary 	./docker/secondary/files/pluggabl/server.run
	cp ./build/pluggabl-exec 		./docker/secondary/files/pluggabl/exec.run
	cd docker && $(MAKE) all
