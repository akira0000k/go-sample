#!/bin/bash
set -x

rm go.mod
rm go.sum
go mod init
go mod tidy

CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o crudapi -ldflags '-s -w' ./cmd/sample-api/main.go
