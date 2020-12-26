@echo off
echo=

echo "需要提前安装配置并安装好minGW"

echo windres
windres -o md5.exe.syso md5.exe.rc

echo build
go build -ldflags="-H windowsgui -w -s"

echo=
pause