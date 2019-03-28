#!/bin/bash

export GO111MODULE=on

go get -d -v ./...
#export GOPATH=$(pwd):$GOPATH
go install -v ./...
