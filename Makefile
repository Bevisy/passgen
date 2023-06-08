.PHONY:
	build
	build-debug
	install
	release
	clean

APP_NAME = passgen
REPO_TAG := $(shell git describe --tags --abbrev=0)

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) main.go

build-debug:
	@echo "Building..."
	@go build -gcflags "all=-N -l" -o bin/$(APP_NAME) main.go

install:
	@echo "Installing..."
	@go install

release:
	@echo "Releasing..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/$(APP_NAME) main.go
	@tar zcf bin/$(APP_NAME)-$(REPO_TAG)-amd64-apple-darwin.tar.gz -C bin $(APP_NAME)
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o bin/$(APP_NAME) main.go
	@tar zcf bin/$(APP_NAME)-$(REPO_TAG)-arm64-apple-darwin.tar.gz -C bin $(APP_NAME)

clean:
	@echo "Cleaning..."
	@rm -rf bin/$(APP_NAME)