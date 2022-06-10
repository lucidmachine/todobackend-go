BIN ?= todobackend-go

.PHONY : all
all: clean build

clean:
	if [ -f $(BIN) ]; then rm $(BIN); fi
	go clean .

build:
	go build -o $(BIN) .

install:
	go install .

test:
	go test ./...

check:
	go vet ./...
	staticcheck ./...

run:
	go run .
