ARG GOVER="1.19.0-alpine3.16"
#ARG GRPCURL_VER="1.8.7"

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

RUN apk update && apk --no-cache add wget
RUN GRPCURL_VER="1.8.7" && wget -q https://github.com/fullstorydev/grpcurl/releases/download/v${GRPCURL_VER}/grpcurl_${GRPCURL_VER}_linux_x86_64.tar.gz && tar -xvzf ./grpcurl_${GRPCURL_VER}_linux_x86_64.tar.gz && chmod a+x ./grpcurl && mv ./grpcurl /usr/local/bin
 
ENV PORT=50051 HEALTH_PORT=5000

VOLUME /app/data
EXPOSE ${PORT} ${HEALTH_PORT}

HEALTHCHECK --interval=30s --timeout=10s --start-period=10s --retries=3 \
	CMD grpcurl -plaintext localhost:${HEALTH_PORT} grpc.healthcheck.Health/Check || exit 1

ENTRYPOINT /app/bggserver --database /app/data/bgg.sqlite --port ${PORT} --healthPort ${HEALTH_PORT} --reflect
