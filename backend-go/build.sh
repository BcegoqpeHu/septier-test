#!/bin/bash

export GO111MODULE=on

go get -d -v ./... \
    && go build -o app server/server.go \
    && chmod +x ./app
