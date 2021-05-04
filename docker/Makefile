.PHONY: all primary secondary lib version

include version.env

version:
	sed -i 's|    version=.*|    version="$(PRIMARY)"|g' ./primary/Dockerfile
	sed -i 's|    version=.*|    version="$(SECONDARY)"|g' ./secondary/Dockerfile
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile

primary:
	docker build -t shanduur/dsim-secondary:$(SECONDARY) -t shanduur/dsim-secondary:latest ./secondary/
	docker push shanduur/dsim-secondary:$(SECONDARY) && docker push shanduur/dsim-secondary:latest

secondary:
	docker build -t shanduur/dsim-primary:$(PRIMARY) -t shanduur/dsim-primary:latest ./primary/
	docker push shanduur/dsim-primary:$(PRIMARY) && docker push shanduur/dsim-primary:latest

all: version primary secondary

lib:
	sed -i 's|    version=.*|    version="$(LIB)"|g' ./lib/Dockerfile
	docker build -t shanduur/dsim-lib:$(LIB) -t shanduur/dsim-lib:latest ./lib/
	docker push shanduur/dsim-lib:$(LIB) && docker push shanduur/dsim-lib:latest