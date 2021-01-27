.PHONY: format test vet

format:
	gofmt -s -w -l .

test:
	go test ./...

vet:
	go vet ./...
