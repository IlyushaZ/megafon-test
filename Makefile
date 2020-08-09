.PHONY: build
build:
	go build -o hasher

.PHONY: run
run:
	go build -o hasher && ./hasher

.PHONY: build-docker
build-docker:
	docker run -p 8087:80 -d ilyushagod/hasher

.PHONY: smoke-test
smoke-test:
	go run smoke-test/main.go

.DEFAULT_GOAL := build-docker