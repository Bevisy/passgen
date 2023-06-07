.PHONY:
	build
	build-debug
	install

APP_NAME=passgen

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) main.go

build-debug:
	@echo "Building..."
	@go build -gcflags "all=-N -l" -o bin/$(APP_NAME) main.go

install:
	@echo "Installing..."
	@go install

clean:
	@echo "Cleaning..."
	@rm -rf bin/$(APP_NAME)