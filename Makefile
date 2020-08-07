.PHONY: run
run:
	go build -o hasher && ./hasher

.PHONY: build-docker
build-docker:
	docker run -p 8087:80 -d ilyushagod/hasher

.DEFAULT_GOAL := build-docker