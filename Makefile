run:
	go build . && ./tcp-chat


build:
	go build .

lint_local:
	golangci-lint run --fix -v

format:
	go fmt ./...

uninstall:
	rm -rf tcp-chat
