GO_ARGS ?=
GOARCH ?= amd64
BUILD = pagesize nginx-property
PLUGIN_NAME = nginx-path-vhosts
GO_PLUGIN_MAKE_TARGET ?= build
BUILD_IMAGE := golang:1.24.2
GO_BUILD_CACHE ?= /tmp/dokku-go-build-cache-$(PLUGIN_NAME)
GO_MOD_CACHE   ?= /tmp/dokku-go-mod-cache-$(PLUGIN_NAME)

.PHONY: build-in-docker build clean src-clean

clean-pagesize:
	rm -rf pagesize

pagesize: clean-pagesize src/pagesize/pagesize.go
	GOARCH=$(GOARCH) cd src/pagesize && go build -ldflags="-s -w" $(GO_ARGS) -o ../../pagesize

clean-nginx-property:
	rm -rf nginx-property

nginx-property: clean-nginx-property src/nginx-property/*.go
	GOARCH=$(GOARCH) cd src/nginx-property && go build -ldflags="-s -w" $(GO_ARGS) -o ../../nginx-property

build: $(BUILD)

build-in-docker: clean
	mkdir -p $(GO_BUILD_CACHE) $(GO_MOD_CACHE)
	docker run --rm \
		-v $(shell pwd):/go/src/nginx-path-vhosts \
		-v $(GO_BUILD_CACHE):/root/.cache \
		-v $(GO_MOD_CACHE):/go/pkg/mod \
		-e GO111MODULE=on \
		-w /go/src/nginx-path-vhosts \
		$(BUILD_IMAGE) \
		bash -c "GO_ARGS='$(GO_ARGS)' CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) GOWORK=off make -j4 $(GO_PLUGIN_MAKE_TARGET)" || exit $$?

clean:
	rm -rf $(BUILD)
	find . -xtype l -delete
