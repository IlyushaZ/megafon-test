.PHONY: build
build:
	go build -o hasher ./cmd/hasher

.PHONY: run
run:
	go build -o hasher ./cmd/hasher && ./hasher

.PHONY: run-container
run-container:
	docker run -p 8087:80 -d ilyushagod/hasher

.PHONY: smoke-test
smoke-test:
	go run ./cmd/smoke-test

.PHONY: unit-tests
unit-tests:
	go test -v ./internal...

.DEFAULT_GOAL := run-container