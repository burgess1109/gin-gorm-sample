.PNONY: run-server

run-server:
	CompileDaemon -log-prefix=false -build="go build -o /usr/bin/go-sample ./main.go" -command="/usr/bin/go-sample"

download:
	go mod download
	go get github.com/githubnemo/CompileDaemon@v1.0.0
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.18.0
	go mod tidy
