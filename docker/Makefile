.PHONY: all primary secondary lib version

include version.env

version:
	sed -i 's|    version=.*|    version="$(PRIMARY)"|g' ./primary/Dockerfile
	sed -i 's|    version=.*|    version="$(SECONDARY)"|g' ./secondary/Dockerfile
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile

primary:
	docker build -t shanduur/dsip-secondary:$(SECONDARY) -t shanduur/dsip-secondary:latest ./secondary/
	docker push shanduur/dsip-secondary:$(SECONDARY) && docker push shanduur/dsip-secondary:latest

secondary:
	docker build -t shanduur/dsip-primary:$(PRIMARY) -t shanduur/dsip-primary:latest ./primary/
	docker push shanduur/dsip-primary:$(PRIMARY) && docker push shanduur/dsip-primary:latest

all: version primary secondary

lib:
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile
	docker build -t shanduur/dsip-lib:$(LIB) -t shanduur/dsip-lib:latest ./lib/
	docker push shanduur/dsip-lib:$(LIB) && docker push shanduur/dsip-lib:latest