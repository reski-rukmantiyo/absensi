#!/bin/bash
sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=linux -e GOARCH=arm -e CGO_ENABLED=1 \
-e CC=arm-linux-gnueabihf-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-armhf-linux-32bit -v source/cmd/main.go

sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=linux -e GOARCH=arm64 -e CGO_ENABLED=1 \
-e CC=aarch64-linux-gnu-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-armhf-linux-64bit -v source/cmd/main.go

sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=windows -e GOARCH=amd64 -e CGO_ENABLED=1 \
-e CC=x86_64-w64-mingw32-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-win64.exe -v source/cmd/main.go