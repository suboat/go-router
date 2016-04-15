#!/usr/bin/env bash
cd websocket
rm -f main
go build main.go server.go client.go
./main
