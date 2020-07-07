#!/bin/bash

version=1.0
comment="kill"

# 交叉编译
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
go build -ldflags  "-X main.BuildVersion=$version -X main.BuildDate=$(date +%F-%T) -X main.BuildComment=$comment"