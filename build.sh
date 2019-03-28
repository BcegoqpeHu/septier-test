#!/bin/bash

export FRONT_DIST=frontend/dist/septier-test-front
export GO_STATIC_PATH=backend-go/public
export JAVA_STATIC_PATH=backend-java/src/main/resources/public

cd frontend
rm -rf dist
./build.sh
cd ..

rm ${JAVA_STATIC_PATH}
cp -R %FRONT_DIST% %JAVA_STATIC_PATH%
cd backend-java
./build.sh
cd ..

rm ${GO_STATIC_PATH}
cp -R %FRONT_DIST% %GO_STATIC_PATH%
cd backend-go
./build.sh
cd ..