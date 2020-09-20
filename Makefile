all: install clean proto gtest

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

clean:
	rm -f pkg/pb/*.pb.go
	rm -f */**/*.tmp.*

gtest:
	go test -cover -race ./...

ptest: proto gtest

install:
	cd $(go env GOPATH)/src/gocv.io/x/gocv && $(MAKE) install