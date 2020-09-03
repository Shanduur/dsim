all: install clean proto gtest

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

clean:
	rm pkg/pb/*.pb.go
	rm */**/*.tmp.*

gtest:
	go test -cover -race ./...

ptest: proto gtest

install:
	go get
	cd $(GOPATH)/src/gocv.io/x/gocv && $(MAKE) install
