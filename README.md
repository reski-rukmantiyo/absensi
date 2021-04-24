
# Quick Start
1. Run the apps
2. Fill all the required info
3. Have fun!!

# Put it as service
In linux, you can try
```
chmod +x absensi
nohup ./absensi >> absensi.log &
```

# Cara Kerja
1. Check untuk file schedule.log
2. kalau tidak ada, create file & schedule.log dan akan buat table adalah jadwal eksekusi untuk hari itu
3. kalau file schedule.log ada dan merupakan tanggal dan jam maka akan eksekusi login atau logout

# Compile to ARM & Windows

Check the result in bin folder

## ARM 32
sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=linux -e GOARCH=arm -e CGO_ENABLED=1 \
-e CC=arm-linux-gnueabihf-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-armhf-linux-32bit -v source/cmd/main.go

## ARM 64
sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=linux -e GOARCH=arm64 -e CGO_ENABLED=1 \
-e CC=aarch64-linux-gnu-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-armhf-linux-64bit -v source/cmd/main.go

## Windows 64bit
sudo docker run -it --rm \
-v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage \
-e GOOS=windows -e GOARCH=amd64 -e CGO_ENABLED=1 \
-e CC=x86_64-w64-mingw32-gcc rrukmantiyo/go-docker-arm-toolchain:latest \
go build -o bin/binary-win64.exe -v source/cmd/main.go

