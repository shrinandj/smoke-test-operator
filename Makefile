GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

OPERATOR_SDK=operator-sdk

all: test build
build:
	$(OPERATOR_SDK) build sjavadekar/smoke-test-operator-base:latest
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)

# Cross compilation
docker: build
	cd docker-build; docker build -t sjavadekar/smoke-test-operator .
