#!/bin/sh
cd /go/src/app
go build -o ./builds web-server.go
go build -o ./builds web-server-with-headers.go
./builds/web-server & ./builds/web-server-with-headers && fg