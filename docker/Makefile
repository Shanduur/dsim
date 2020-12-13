.PHONY: all primary secondary lib version

include version.env

version:
	sed -i 's|    version=.*|    version="$(PRIMARY)"|g' ./primary/Dockerfile
	sed -i 's|    version=.*|    version="$(SECONDARY)"|g' ./secondary/Dockerfile
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile

primary:
	docker build -t shanduur/pluggabl-secondary:$(SECONDARY) -t shanduur/pluggabl-secondary:latest ./secondary/
	docker push shanduur/pluggabl-secondary:$(SECONDARY) && docker push shanduur/pluggabl-secondary:latest

secondary:
	docker build -t shanduur/pluggabl-primary:$(PRIMARY) -t shanduur/pluggabl-primary:latest ./primary/
	docker push shanduur/pluggabl-primary:$(PRIMARY) && docker push shanduur/pluggabl-primary:latest

all: version primary secondary

lib:
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile
	docker build -t shanduur/pluggabl-lib:$(LIB) -t shanduur/pluggabl-lib:latest ./lib/
	docker push shanduur/pluggabl-lib:$(LIB) && docker push shanduur/pluggabl-lib:latest