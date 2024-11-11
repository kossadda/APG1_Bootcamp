GO := go
GOFILES := $(shell find . -name '*.go')
EXE := program

all: build run

build:
	@if [ ! -f go.mod ]; then \
		$(GO) mod init main; \
	fi
	$(GO) build -o $(EXE) main.go

run:
	./$(EXE)

clean:
	rm -f go.mod
	rm -f $(EXE)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
