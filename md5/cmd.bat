@echo off
echo=

echo "���Ȱ�װ��dev C++����minGW"

echo windresing...
windres -o md5.exe.syso md5.exe.rc

echo building...
go build -ldflags="-H windowsgui -w -s"

echo=
pause