run:
	go build . && ./tcp-chat

lint_local:
	golangci-lint run --fix -v

format:
	go fmt ./...
