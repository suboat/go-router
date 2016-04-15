#!/usr/bin/env bash
cd http
rm -f main
go build main.go server.go client.go
./main
