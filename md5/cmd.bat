@echo off
echo=

:: 1. minGW����bug��������޷�ִ�У����о����������޸�

:: echo "���Ȱ�װ��dev C++����minGW"

:: echo windresing...
:: windres -o md5.exe.syso md5.exe.rc

:: echo building...
:: go build -ldflags="-H windowsgui -w -s"

:: 2.rsrc

echo "��Ҫ�Ȱ�װ��������rsrc�����go get github.com/akavel/rsrc" 

echo rsrc manifesting...
rsrc -manifest md5.exe.manifest -o md5.exe.syso -ico md5.ico

echo building...
go build

echo=
pause