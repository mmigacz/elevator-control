build:
	go build

test:
	go test

clean:
	go fmt ../...
	go vet ../...
	go clean



