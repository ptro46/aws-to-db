#!/bin/bash

export GOPATH=`pwd`
go env -w GO111MODULE=off
go get github.com/fatih/color
go get github.com/lib/pq
go build aws-to-db
