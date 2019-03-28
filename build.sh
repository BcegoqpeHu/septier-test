#!/bin/bash

set -e

STATIC_JAVA_PATH=backend-java/src/main/resources/public
STATIC_GO_PATH=backend-go/src/main/resources/public

cd frontend && build.sh && cd ..

cd ..\backend-go
build.cmd

# Сборка Java - бэкенда

cp frontend/dist backend-java/src/main/resources/public
cd ..\backend-java
build.cmd