VERSION=0.1.0
TARGETS_NOVENDOR=$(shell glide novendor)

all: bin/imshell

bundle:
	glide install

bin/imshell: cmd/imshell/imshell.go imshell/*.go
	go build -o bin/imshell cmd/imshell/imshell.go

fmt:
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

check:
	go test $(TARGETS_NOVENDOR)

clean:
	rm -rf bin/*
