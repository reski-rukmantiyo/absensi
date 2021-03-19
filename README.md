
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

4. Run absensi.exe or absensi (in linux)
5. Just let your apps running all day

# Cara Kerja
1. Check untuk file schedule.log
2. kalau tidak ada, create file & schedule.log dan akan buat table adalah jadwal eksekusi untuk hari itu
3. kalau file schedule.log ada dan merupakan tanggal dan jam maka akan eksekusi login atau logout

