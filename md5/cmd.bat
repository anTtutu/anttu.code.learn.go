@echo off
echo=

:: 1. minGW，有bug，打包后无法执行，等研究合适了再修复

:: echo "请先安装好dev C++工具minGW"

:: echo windresing...
:: windres -o md5.exe.syso md5.exe.rc

:: echo building...
:: go build -ldflags="-H windowsgui -w -s"

:: 2.rsrc

echo "需要先安装工具命令rsrc，命令：go get github.com/akavel/rsrc" 

echo rsrc manifesting...
rsrc -manifest md5.exe.manifest -o md5.exe.syso -ico md5.ico

echo building...
go build

echo=
pause