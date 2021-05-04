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
	[ ! -d /opt/dsim/ ] && sudo mkdir /opt/dsim/ || echo ok

	sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/primary.run   /opt/dsim/primary.run
	sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/client.run    /opt/dsim/client.run
	[ -e ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run ] && \
		sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run /opt/dsim/secondary.run || echo ok
	[ -e ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run ] && \
		sudo cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run      /opt/dsim/exec.run  || echo ok

	sudo cp ./scripts/dsim.sh      /opt/dsim/dsim.sh

	sudo ln -sf /opt/dsim/dsim.sh /usr/bin/dsim
	
	sudo chmod +x /opt/dsim/*

	[ ! -d /etc/dsim/ ] && sudo mkdir /etc/dsim/ || echo ok
	sudo cp ./config/config_primary.json /etc/dsim/
	sudo cp ./config/config_secondary.json /etc/dsim/
	sudo cp ./config/config_db.json /etc/dsim/

	[ ! -d ~/.config/ ] && sudo mkdir ~/.config/ || echo ok
	[ ! -d ~/.config/dsim.d/ ] && sudo mkdir ~/.config/dsim.d/ || echo ok
	sudo cp ./config/config_client.json ~/.config/dsim.d/

test:
	go test -cover -race ./...
	
docker: build_all
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/primary.run   ./docker/primary/files/dsim/server.run
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/secondary.run ./docker/secondary/files/dsim/server.run
	cp ./build/$(shell go env GOOS)/$(shell go env GOARCH)/exec.run      ./docker/secondary/files/dsim/exec.run
	cd docker && $(MAKE) all

build_multiarch:
	echo "This supports only 64-bit architectures. Also you are not building dsim/exec." ; \
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