PACKAGES=$(shell go list ./... | grep -v 'tests')

### Tools needed for development
devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/tetafro/godot/cmd/godot@latest
	go install mvdan.cc/gofumpt@latest

## build and release
build:
	go build -o bin/taar

releaser:
	goreleaser release

cp:
	sudo cp bin/taar /usr/bin 

### Formatting, linting, and vetting
fmt:
	gofumpt -l -w .
	godot -w .

check:
	golangci-lint run --build-tags "${BUILD_TAG}" --timeout=20m0s