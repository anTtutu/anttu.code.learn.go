@echo off
Title build...                                              
Color 0A

REM 声明采用UTF-8编码
chcp 65001

:menu
echo =============================================
echo.
echo No.1-选择用minGW工具打包
echo.
echo No.2-选择用rsrc工具打包
echo.
echo =============================================

set /p user_input=请输入你的选择:

if %user_input% equ 1 goto 1

if %user_input% equ 2 goto 2

if %user_input% == '3' exit

pause

:: 1. minGW
:1
echo "请先安装好dev C++工具minGW"

echo windresing...
windres -o md5.exe.syso md5.exe.rc

echo building...
go build
goto end

:: 2.rsrc
:2
echo "需要先安装工具命令rsrc，命令：go get github.com/akavel/rsrc" 

echo rsrc manifesting...
rsrc -manifest md5.exe.manifest -o md5.exe.syso -ico md5.ico

echo building...
go build
goto end

:end
echo complete