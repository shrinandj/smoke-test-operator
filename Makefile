GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

OPERATOR_SDK=operator-sdk

all: test build
build:
	$(OPERATOR_SDK) build ${IMAGE_NAMESPACE}/smoke-test-operator-base:latest
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)

docker: build
	cd docker-build; docker build -t ${IMAGE_NAMESPACE}/smoke-test-operator .

docker-push: docker
	docker push ${IMAGE_NAMESPACE}/smoke-test-operator
