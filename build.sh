#! /bin/sh
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o todo main.go
