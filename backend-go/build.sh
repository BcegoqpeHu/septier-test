#!/bin/bash

go get ./...
export GOPATH=$(pwd):$GOPATH

go build src/septier_test_backend/server.go