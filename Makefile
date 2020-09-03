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

gatest:
	go test $(go list ./... | grep -v /compute/)

install:
	cd $(go env GOPATH)/src/gocv.io/x/gocv && $(MAKE) install