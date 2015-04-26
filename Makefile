install:
	go install

test:
	go test

clean:
	go fmt ../...
	go vet ../...
	go clean



