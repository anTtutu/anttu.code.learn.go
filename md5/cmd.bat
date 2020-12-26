@echo off
echo=

echo "请先安装好dev C++工具minGW"

echo windresing...
windres -o md5.exe.syso md5.exe.rc

echo building...
go build -ldflags="-H windowsgui -w -s"

echo=
pause