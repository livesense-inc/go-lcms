MAKEFLAGS += --warn-undefined-variables

test:
	@go clean -testcache
	@go test -race ./...

lint:
	@go vet ./...

run: main.go
	@go run $^
