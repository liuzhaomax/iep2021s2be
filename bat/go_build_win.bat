@cd "%~dp0%"
@cd..
@cd src\main
:@cls
:@cmd
set GO111MODULE=on
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
@REM set CC=x86_64-w64-mingw32-gcc
@REM set CXX=x86_64-w64-mingw32-g++
go mod tidy
go build -o ../../bin/