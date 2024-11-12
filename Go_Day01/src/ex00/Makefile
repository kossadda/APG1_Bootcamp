GO := go
EXE := readDB
ARG := 

all: build run

build:
	@if [ ! -f go.mod ]; then \
		$(GO) mod init main; \
	fi
	@$(GO) build -o $(EXE) main.go

run:
	@./$(EXE) $(ARG)

clean:
	rm -f go.mod
	rm -f $(EXE)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
