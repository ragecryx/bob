# Project parameters
PACKAGE=github.com/ragecryx/bob

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bob_builder

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/stretchr/testify/assert
		$(GOGET) github.com/sirupsen/logrus
		$(GOGET) gopkg.in/yaml.v2
		$(GOGET) gopkg.in/src-d/go-git.v4
		$(GOGET) github.com/gorilla/mux
		$(GOGET) github.com/yosssi/ace


# Cross compilation
build-windows:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_win_amd64
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux_amd64
build-osx:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_osx_amd64

docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/$(PACKAGE) golang:latest go build -o "$(BINARY_UNIX)" -v
