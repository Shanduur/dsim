GOCV_VERSION=v0.25.0 

.PHONY: build build_all build_multiarch prebuild install clean proto test docker

all: prebuild clean proto build

prebuild:
	go mod download
	cd $(shell go env GOPATH)/pkg/mod/gocv.io/x/gocv\@$(GOCV_VERSION) && $(MAKE) install

clean:
	rm -f ./pb/*.pb.go
	rm -rf ./build/*

proto:
	protoc \
	--proto_path=proto \
	--go_out=plugins=grpc:. \
	proto/*.proto

build:
	go build -o ./build/$(shell go env GOOS)/$(shell go env GOARCH)/primary.run   ./cmd/server/primary 
	go build -o ./build/$(shell go env GOOS)/$(shell go env GOARCH)/client.run    ./cmd/client

build_all: build
	go build -o ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run ./cmd/server/secondary 
	go build -o ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run      ./cmd/exec 

install:
	[ ! -d /opt/pluggabl/ ] && sudo mkdir /opt/pluggabl/ || echo ok

	sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/primary.run   /opt/pluggabl/primary.run
	sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/client.run    /opt/pluggabl/client.run
	[ -e ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run ] && \
		sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run /opt/pluggabl/secondary.run || echo ok
	[ -e ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run ] && \
		sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run      /opt/pluggabl/exec.run  || echo ok

	sudo cp ./scripts/pluggabl.sh      /opt/pluggabl/pluggabl.sh

	sudo ln -sf /opt/pluggabl/pluggabl.sh /usr/bin/pluggabl
	
	sudo chmod +x /opt/pluggabl/*

	[ ! -d /etc/pluggabl/ ] && sudo mkdir /etc/pluggabl/ || echo ok
	sudo cp ./config/config_primary.json /etc/pluggabl/
	sudo cp ./config/config_secondary.json /etc/pluggabl/
	sudo cp ./config/config_db.json /etc/pluggabl/

	[ ! -d ~/.config/ ] && sudo mkdir ~/.config/ || echo ok
	[ ! -d ~/.config/pluggabl.d/ ] && sudo mkdir ~/.config/pluggabl.d/ || echo ok
	sudo cp ./config/config_client.json ~/.config/pluggabl.d/

test:
	go test -cover -race ./...
	
docker: build_all
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/primary.run   ./docker/primary/files/pluggabl/server.run
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run ./docker/secondary/files/pluggabl/server.run
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run      ./docker/secondary/files/pluggabl/exec.run
	cd docker && $(MAKE) all

build_multiarch:
	echo "This supports only 64-bit architectures. Also you are not building pluggabl/exec." ; \
	for i in 0 ; do \
		for os in windows ; do \
			for arch in amd64 ; do \
				echo $$os $$arch ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/primary.exe   ./cmd/server/primary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/secondary.exe ./cmd/server/secondary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/client.exe    ./cmd/client ; \
			done ; \
		done ; \
		for os in linux darwin ; do \
			for arch in amd64; do \
				echo $$os $$arch ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/primary.run   ./cmd/server/primary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/secondary.run ./cmd/server/secondary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/client.run    ./cmd/client ; \
			done ; \
		done ; \
		for os in linux ; do \
			for arch in arm64 ppc64 ppc64le riscv64 ; do \
				echo $$os $$arch ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/primary.run   ./cmd/server/primary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/secondary.run ./cmd/server/secondary ; \
				GOOS=$$os GOARCH=$$arch go build -o ./build/$$os/$$arch/client.run    ./cmd/client ; \
			done ; \
		done \
	done ; \