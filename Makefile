# https://www.xiexianbin.cn/program/tools/2016-01-09-makefile/index.html
.PHONY: all test clean

GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

help:  ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: clean test  ## Build all
test:  ## run test
	$(GOTEST) -v ./...
clean: ## run clean bin files
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
