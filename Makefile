SHELL := /bin/bash
VERSION := 0.1.0

.PHONY: version
version:
	@echo "whroom: $(VERSION)"


.PHONY: build
build:
	go get ./...
	go build

.PHONY: brew-update
brew-update:
	bash .circleci/scripts/entrypoint.bash $(VERSION)
