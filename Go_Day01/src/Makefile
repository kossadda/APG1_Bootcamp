GO := go
EXE00 := readDB
EXE01 := compareDB
EXE02 := compareFS
ARG := 

deb:
	@go build -o $(EXE02) ex02/main.go
	@./$(EXE02)

all: build run

build_ex00:
	@$(GO) build -o $(EXE00) ex00/main.go

build_ex01:
	@$(GO) build -o $(EXE01) ex01/main.go

build_ex02:
	@$(GO) build -o $(EXE02) ex02/main.go

clean:
	rm -f $(EXE00)
	rm -f $(EXE01)
	rm -f $(EXE02)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy
