run:
	@make build
	@./bin/accesslog-replayer

build:
	@go clean
	@go build -o bin/accesslog-replayer ./cmd/
