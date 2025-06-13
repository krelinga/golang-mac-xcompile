#! /usr/bin/bash

export GOOS=linux
export GOARCH=amd64

go build -o /nas/dev/golang-linux-xcompile-hello-world .
