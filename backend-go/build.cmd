set GO111MODULE=on
go get -d -v ./... && go build -o app.exe server/server.go