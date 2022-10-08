#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/routecheck main.go
GOOS=windows GOARCH=amd64 go build -o $GOPATH/bin/routecheck.exe main.go