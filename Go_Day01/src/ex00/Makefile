GO := go
EXE := readDB
ARG := 

all: build run

build:
	@$(GO) build -o $(EXE) main.go

run:
	@./$(EXE) $(ARG)

clean:
	rm -f $(EXE)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
