PROTO_FILES=$(shell find ./proto/p2p -type f -name "*.proto" | sed 's/\/proto\/p2p//g')

build:
	go build -o bin/oasvlfy .

.PHONY: proto
proto:
	protoc --go_out=./proto/p2p/v2  \
	./proto/p2p/v2/*.proto

fmt:
	go fmt ./...

fmtproto:
	clang-format -i ./proto/p2p/**/*.proto

test:
	go test -v ./...