build:
	go build -o bin/oasvlfy .

fmt:
	go fmt ./...

test:
	go test -v ./...