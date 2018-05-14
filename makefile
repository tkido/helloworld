# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v
GOGET=$(GOCMD) get -u
BINARY_NAME=mmm.exe
BINARY_UNIX=$(BINARY_NAME)_unix

all: test clean asset build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) .
	$(GOTEST) ./core/godfather
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
asset:
	go-assets-builder -o=assets.go -p=main assets/
deps:
	$(GOGET) github.com/hajimehoshi/ebiten
	$(GOGET) github.com/golang/freetype/truetype
	$(GOGET) github.com/jessevdk/go-assets
	$(GOGET) github.com/jessevdk/go-assets-builder
	$(GOGET) golang.org/x/image/font

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v