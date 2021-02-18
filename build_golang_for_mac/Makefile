
TAG = $(shell git describe --always --tags --abbrev=0 | tr -d "[v\r\n]")
COMMIT = $(shell git rev-parse --short HEAD| tr -d "[ \r\n\']")
VERSION :=v$(TAG)-$(COMMIT)
BUILD_TIME := $(shell date +%Y%m%d-%H%M%S)

FLAGS = -gcflags "-N -l" -race
GOFLAGS = -i
LD_FLAGS = "-extldflags -static -w -s -X main.version=v$(TAG)-$(COMMIT) -X 'main.buildTime=$(BUILD_TIME)'"

PACKAGE = ./
BIN = $(project)

.SUFFIXES:

all: svc

svc:
	@go build $(GOFLAGS) -ldflags=$(LD_FLAGS) -o ./$(BIN) $(PACKAGE)

build:
	./build.sh

check:
	gofmt -w -s -d .
	go vet "./..."

docker:
	./docker_build.sh

clean:
	@rm -f $(BIN)
	@rm -f $(BIN).tar.gz

.PHONY: svc
