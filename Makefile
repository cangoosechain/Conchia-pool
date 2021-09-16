run:
	go run .

build:
	go build .

build-and-test:
	go build .
	go test ./...

build-linux:
	GOOS=linux GOARCH=amd64 go build
