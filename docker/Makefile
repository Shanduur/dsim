.PHONY: all build push lib version

include version.env

version:
	sed -i 's|.*version=.*|version="$(VERSION)"|g' ./primary/Dockerfile
	sed -i 's|.*version=.*|version="$(VERSION)"|g' ./secondary/Dockerfile
	sed -i 's|.*version=.*|version="$(LIB)"|g' ./lib/Dockerfile

build:
	docker build -t shanduur/pluggabl-primary:$(VERSION) 	./primary/
	docker build -t shanduur/pluggabl-secondary:$(VERSION) 	./secondary/

push:
	docker push shanduur/pluggabl-primary:$(VERSION)
	docker push shanduur/pluggabl-secondary:$(VERSION)

all: version build push

lib:
	docker build -t shanduur/pluggabl-lib:$(LIB)
	docker push 	shanduur/pluggabl-lib:$(LIB)