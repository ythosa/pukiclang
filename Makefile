.PHONY: run
run:
	go run ./src/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v ./src/...

.PHONY: pipeline
pipeline:
	make test && make

.DEFAULT_GOAL := run
