FROM golang:1.12

WORKDIR $GOPATH/src/septier-test

COPY . .
RUN ./build.sh

EXPOSE 8080
CMD ["./app"]