@cd "%~dp0%"
@cd..
@cd src\main
:@cls
:@cmd
set GO111MODULE=on
set CGO_ENABLED=1
set GOOS=linux
set GOARCH=amd64
@REM set CC=x86_64-linux-musl-gcc
@REM set CXX=x86_64-linux-musl-g++
go mod tidy
go build -o ../../bin/
pause