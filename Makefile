# includes the root level Makefile
-include ../Makefile

OMD_COMPONENT := sample-api
SOURCE := ./...
CONTAINER_IMAGE ?= $(shell basename $(shell make github-repo))/$(OMD_COMPONENT):latest

.DEFAULT_GOAL := build

# Go Targets
build:
	go build -v -o $(OMD_COMPONENT)
.PHONY: build

lint:
	go vet $(SOURCE)
	test -z $(shell go fmt $(SOURCE))
.PHONY: lint

unit-test:
	go test -coverprofile=coverage.out $(SOURCE) -count=1
.PHONY: unit-test

.PHONY: mocks
mocks: ## create mocks
	GO111MODULE=on mockery --dir=controller --name=Usecase --case=snake --output=controller/mocks --outpkg=mocks
	GO111MODULE=on mockery --dir=router --name=Controller --case=snake --output=router/mocks --outpkg=mocks
	GO111MODULE=on mockery --dir=usecase --name=Service --case=snake --output=usecase/mocks --outpkg=mocks
