GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

OPERATOR_SDK=operator-sdk

IMAGE_VERSION=latest
ifdef RELEASE_VERSION
IMAGE_VERSION=${RELEASE_VERSION}
endif

all: test build
build:
	$(OPERATOR_SDK) build ${IMAGE_NAMESPACE}/smoke-test-operator-base:${IMAGE_VERSION}
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)

docker: build
	cd docker-build; docker build -t ${IMAGE_NAMESPACE}/smoke-test-operator:${IMAGE_VERSION} .

docker-push: docker
	docker push ${IMAGE_NAMESPACE}/smoke-test-operator:${IMAGE_VERSION}
