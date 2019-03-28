#!/bin/bash

./build.sh

docker build -t sotnikov/septier-test-go:latest ./backend-go/
docker build -t sotnikov/septier-test-java:latest ./backend-java/

docker push sotnikov/septier-test-go:latest
docker push sotnikov/septier-test-java:latest
