build:
	go build -o ./bin/file-checksum

run: build
	./bin/file-checksum

test:
	go test -v ./..