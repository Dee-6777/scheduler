export GO111MODULE=on
export CGO_ENABLED=0

# Constants:
BINARY=scheduler

.PHONY: build
build:
	go build

.PHONY: mod
mod:
	go mod tidy

.PHONY: test
test: mod vendor
	go test ./tests -v

.PHONY: clean
clean:
	rm ${BINARY}