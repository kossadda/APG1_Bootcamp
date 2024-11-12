GO := go
EXE00 := readDB
EXE01 := compareDB
ARG := 

deb:
	@go build -o $(EXE01) ex01/main.go
	@./$(EXE01) --old=test/original_database.xml --new=test/stolen_database.json

all: build run

build_ex00:
	@$(GO) build -o $(EXE00) ex00/main.go

build_ex01:
	@$(GO) build -o $(EXE01) ex01/main.go

clean:
	rm -f $(EXE00)
	rm -f $(EXE01)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
