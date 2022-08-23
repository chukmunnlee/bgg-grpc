.DEFAULT_GOAL := build

PROTOC=protoc

gen-grpc: grpc/bgg_grpc.pb.go grpc/bgg.pb.go
	$(PROTOC) --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		grpc/bgg.proto

build: linux windows darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/bggserver-linux-amd64 .

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/bggserver-windows-amd64 .

darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/bggserver-darwin-amd64 .

clean:
	rm -rf bin

