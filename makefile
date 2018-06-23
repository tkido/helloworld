# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v
GOGET=$(GOCMD) get
BINARY_NAME=mmm.exe
BINARY_UNIX=$(BINARY_NAME)_unix

PKG_LIFE=exsamples/life
PKG_COLLISION=exsamples/collision

all: test clean asset build
run: clean build execute
deps:
	$(GOGET) github.com/atotto/clipboard
	$(GOGET) github.com/hajimehoshi/ebiten
	$(GOGET) github.com/golang/freetype/truetype
	$(GOGET) github.com/jessevdk/go-assets
	$(GOGET) github.com/jessevdk/go-assets-builder
	$(GOGET) golang.org/x/image/font
asset:
	go-assets-builder -p=assets -o=assets/assets.go -v=FileSystem -s=/_assets/ _assets/
	# go-assets-builder -p=main -o=assets.go assets/
	# go-assets-builder -p=main -o=$(PKG_LIFE)/assets.go -s=/$(PKG_LIFE) $(PKG_LIFE)/assets/
	# go-assets-builder -p=main -o=$(PKG_COLLISION)/assets.go -s=/$(PKG_COLLISION) $(PKG_COLLISION)/assets/
test:
	$(GOTEST) .
	$(GOTEST) ./core/godfather
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
execute:
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v