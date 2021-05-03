.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "./bin/goget"  -v ./main.go

.PHONY: run
run:
	bin/goget http://speedtest.singapore.linode.com/100MB-singapore.bin test.bin

DEFAULT_COAL := build