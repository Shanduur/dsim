.PHONY: all build push lib

LIB=0.0.1-init
VER=0.0.1-init

build:
	docker build -t shanduur/pluggabl-primary:$(VER) 	./primary/
	docker build -t shanduur/pluggabl-secondary:$(VER) 	./secondary/

push:
	docker push shanduur/pluggabl-primary:$(LIB)
	docker push shanduur/pluggabl-secondary:$(LIB)

all: build push

lib:
	docker build -t shanduur/pluggabl-lib:$(LIB)
	docker push 	shanduur/pluggabl-lib:$(LIB)