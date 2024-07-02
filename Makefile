build:
	go build -o bin/oasvlfy .

.PHONY: proto
proto:
	protoc --go_out=./proto/p2p/v1  \
	./proto/p2p/v1/*.proto

fmt:
	go fmt ./...

fmtproto:
	clang-format -i ./proto/p2p/**/*.proto

test:
	go test -v ./...