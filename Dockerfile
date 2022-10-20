ARG GOVER="1.19.0-alpine3.16"

FROM golang:${GOVER} AS builder

LABEL org.opencontainers.image.source=https://github.com/chukmunnlee/bgg-grpc

WORKDIR /go/src/github.com/chukmunnlee/bgg-grpc

RUN apk add --no-cache make build-base

ADD go.mod .
ADD go.sum .
ADD bggserver.go .
ADD data/bggdb.go data/bggdb.go
ADD grpc grpc

RUN go mod download
#RUN CGO_ENABLED=0 go build -o bggserver .
RUN go build -o bggserver .

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /go/src/github.com/chukmunnlee/bgg-grpc/bggserver /app/bggserver

ADD data/bgg.sqlite data/bgg.sqlite

ENV PORT=50051

VOLUME /app/data
EXPOSE ${PORT}

ENTRYPOINT /app/bggserver --database /app/data/bgg.sqlite --port ${PORT} --reflect
