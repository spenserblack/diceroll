.PHONY: clean format lint test vet

clean:
	rm diceroll

diceroll:
	go build -o diceroll ./cmd/diceroll/

format:
	gofmt -s -w -l .

lint: vet
	golint ./...

test:
	go test ./...

vet:
	go vet ./...
