all: clean proto gtest

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
	ls -la /
	tree /
	cd $(go env GOPATH)/src/gocv.io/x/gocv && $(MAKE) install
