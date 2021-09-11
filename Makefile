# include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
#GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
GOSOURCEFILES="source/cmd/main.go"
TAG:=$(shell git describe --tags --abbrev=0 --always)

build:
	go build -o $(GOBIN)/$(PROJECTNAME) -v $(GOSOURCEFILES) 

build-multi:
	## ARM 32
	sudo docker run -it --rm \
	-v "$(GOBASE)":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
	-e GOOS=linux -e GOARCH=arm -e CGO_ENABLED=1 \
	-e CC=arm-linux-gnueabihf-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
	go build -o bin/absensi-linux-arm32 -v source/cmd/main.go

	# ## ARM 64
	sudo docker run -it --rm \
	-v "$(GOBASE)":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
	-e GOOS=linux -e GOARCH=arm64 -e CGO_ENABLED=1 \
	-e CC=aarch64-linux-gnu-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
	go build -o bin/absensi-linux-arm64 -v source/cmd/main.go

	## Windows 64bit
	sudo docker run -it --rm \
	-v "$(GOBASE)":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
	-e GOOS=windows -e GOARCH=amd64 -e CGO_ENABLED=1 \
	-e CC=x86_64-w64-mingw32-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
	go build -o bin/absensi-win64.exe -v source/cmd/main.go

	## Linux 64bit
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(GOBIN)/$(PROJECTNAME) -v $(GOSOURCEFILES) 

build-image:
	echo "Make docker image and push to docker hub with default is for AMD64"
	sudo docker build -t absensi:latest -f docker/dockerfile-linux-amd64 .
	sudo docker build -t absensi:amd64-v$(TAG) -f docker/dockerfile-linux-amd64 .
	# sudo docker build -t absensi:win64-v$(TAG) -f docker/dockerfile-windows-amd64 .
	# sudo docker build -t absensi:linux-arm32-v$(TAG) -f docker/dockerfile-linux-arm32 .
	# sudo docker build -t absensi:linux-arm64-v$(TAG) -f docker/dockerfile-linux-arm64 .