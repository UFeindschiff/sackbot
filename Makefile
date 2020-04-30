GPATH=$(shell pwd):$(shell pwd)/vendor
export PATH := $(PATH):$(shell pwd)/bin:$(GPATH)/bin

.PHONY: all sackbot

all: sackbot

sackbot: bin/sackbot

clean:
	rm -rf bin
	go clean

bin/sackbot: *.go
	mkdir -p bin
	go build -o $@


