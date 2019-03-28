FROM golang:1.12
WORKDIR $GOPATH/src/septier-test
COPY . .
RUN ./build.sh
EXPOSE 8080
RUN ls $GOPATH/bin
CMD ["backend-go"]