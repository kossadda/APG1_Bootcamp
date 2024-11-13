GO := go
EXE := program

all: build run

build:
	$(GO) build -o $(EXE) main.go

run:
	./$(EXE)

clean:
	rm -f $(EXE)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
