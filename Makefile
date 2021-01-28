.PHONY: format lint test vet

format:
	gofmt -s -w -l .

lint: vet
	golint ./...

test:
	go test ./...

vet:
	go vet ./...
