CC=go

BUILD=build
BUILD_FLAGS=-o

FMT=gofmt
FMT_FLAGS=-s -w

OUT=main

all: run

run:
	$(CC) run main.go

build:
	$(CC) $(BUILD) $(BUILD_FLAGS) $(OUT) main.go

fmt:
	$(FMT) $(FMT_FLAGS) .
