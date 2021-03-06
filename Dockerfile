FROM golang:1.14-alpine AS build
ENV CGO_ENABLED=0
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o /bin/hasher cmd/hasher/main.go

FROM alpine:latest
COPY --from=build /bin/hasher /bin/hasher

ENTRYPOINT ["/bin/hasher"]