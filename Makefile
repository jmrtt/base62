.PHONY: help 
help: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: dependencies
dependencies: ## Download dependencies required by application
	go mod download

.PHONY: lint
lint: ## Perform linting to the project
	golangci-lint run -v ./...

.PHONY: test
test: ## Run tests for the application
	go test -cover ./...

.PHONY: build-linux
build-linux: ## Build application for linux
	GOOS=linux GOARCH=amd64 go build -o bin/base62-linux

.PHONY: build-darwin
build-darwin: ## Build application for macos
	GOOS=darwin GOARCH=amd64 go build -o bin/base62-darwin

.PHONY: build-windows
build-windows: ## Build application for windows
	GOOS=windows GOARCH=amd64 go build -o bin/base62-windows.exe

.PHONY: build
build: ## Build for all platforms
build: build-linux build-darwin build-windows