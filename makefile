install:
	go get -u -d gocv.io/x/gocv
	cd $(GOPATH)/src/gocv.io/x/gocv && $(MAKE) install

