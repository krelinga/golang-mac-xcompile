#! /usr/bin/bash

export GOOS=darwin
export GOARCH=arm64

go build -o /nas/dev/golang-mac-xcompile-hello-world .
