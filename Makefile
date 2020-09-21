# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

GOFMT=gofmt
GOLINT=./bin/golangci-lint
BINARY_NAME=yamllint-checkstyle
CMD_FOLDER=cmd
COVERAGE_FOLDER=.coverage
RELEASE_VERSION?=unset
GOLANGCI_VERSION=v1.31.0

all: generate build lint test

clean:
	$(GOCMD) clean -i .

clean-all: clean ## remove all generated artifacts and clean all build artifacts
	rm -rf docs vendor $(COVERAGE_FOLDER) $(BINARY_NAME)
	find . -name "*_gen.go" -delete # remove generated files
	find . -name "*_generated.go" -delete # remove generated files

vendor: ## update golang dependencies
	go mod vendor

test: vendor
	$(GOTEST) -short -mod=vendor ./...

build: vendor
	$(GOBUILD) -ldflags "-X main.version=$(RELEASE_VERSION)" -mod=vendor -o $(BINARY_NAME) main.go

coverage:
	mkdir -p .coverage/
	$(GOTEST) -short -mod=vendor -coverprofile=.coverage/profile.cov.tmp ./...
	cat .coverage/profile.cov.tmp | grep -v "_gen.go"> .coverage/profile.cov
ifeq ($(CI), true)
	GO111MODULE=off $(GOGET) github.com/mattn/goveralls
	goveralls -coverprofile=.coverage/profile.cov -service=travis-ci
endif

lint:
ifeq ($(CI), true)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(GOLANGCI_VERSION)
endif
	$(GOLINT) run --config ./.golangci.yml

# Setup your environment
setup: vendor generate
