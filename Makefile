run:
	go run main.go

lint_local:
	golangci-lint run --fix -v

format:
	go fmt ./...
