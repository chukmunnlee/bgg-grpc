.DEFAULT_GOAL := build

PROTOC=protoc

grpc: 
	$(PROTOC) --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		grpc/server/bgg.proto
	$(PROTOC) --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		grpc/healthcheck/healthcheck.proto

gateway: 
	$(PROTOC) -I . --grpc-gateway_out=. \
		--grpc-gateway_opt=paths=source_relative \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=generate_unbound_methods=true \
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

