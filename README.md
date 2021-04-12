
# Quick Start
1. Copy file .env from examples/.env to same path with absensi.exe or absensi
2. Copy your picture to upload to same path with absensi.exe or absensi
3. Fill contents of .env files. 
```
BaseURL=https://myapps.lintasarta.net/api
UserName=
Password=
Picture=
Longitude=
Lattitude=
Description=
Region=Asia/Jakarta
```
Tips: Dont change for BaseURL and Region, Picture only filename with extension

4. Download and Run absensi.exe or absensi (in linux) from [Release Page](https://github.com/reski-rukmantiyo/absensi/releases)
5. Just let your apps running all day

# Cara Kerja
1. Check untuk file schedule.log
2. kalau tidak ada, create file & schedule.log dan akan buat table adalah jadwal eksekusi untuk hari itu
3. kalau file schedule.log ada dan merupakan tanggal dan jam maka akan eksekusi login atau logout

# Compile to ARM

## ARM 32
sudo docker run -it --rm -v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage -e GOOS=linux -e GOARCH=arm -e CGO_ENABLED=1 -e CC=arm-linux-gnueabihf-gcc rrukmantiyo/go-docker-arm-toolchain:latest go build -o linux-arm32 -v source/cmd/main.go

## ARM 64
sudo docker run -it --rm -v "$PWD":/go/src/myrepo/mypackage -w /go/src/myrepo/mypackage -e GOOS=linux -e GOARCH=arm64 -e CGO_ENABLED=1 -e CC=aarch64-linux-gnu-gcc rrukmantiyo/go-docker-arm-toolchain:latest go build -o linux-arm64 -v source/cmd/main.go
