SHELL = /bin/bash

# Get the build dependencies
build_dep:
	go get -t ./...
	go get -u github.com/kisielk/errcheck
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint