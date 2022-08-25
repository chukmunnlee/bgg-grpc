ARG GOVER="1.19.0"

FROM golang:${GOVER} AS builder

LABEL org.opencontainers.image.source=https://github.com/chukmunnlee/bgg-grpc

WORKDIR /go/src/github.com/chukmunnlee/bgg-grpc

ADD go.mod .
ADD go.sum .
ADD bggserver.go .
ADD data/bggdb.go data/bggdb.go
ADD grpc grpc

RUN go mod download
#RUN CGO_ENABLED=0 go build -o bggserver .
RUN go build -o bggserver .

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /go/src/github.com/chukmunnlee/bgg-grpc/bggserver /app/bggserver

ADD data/bgg.sqlite data/bgg.sqlite

VOLUME /app/data

CMD [ "/app/bggserver", "--database", "/app/data/bgg.sqlite", "--port", "50051", "--reflect" ]
