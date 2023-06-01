# ================= #
# goecm build entry #
# ================= #


OUTPUT_DIR := /proj/goecm
ROOT_DIR := .
BLOCKER_TOOLS ?= golingci-lint statik
DEPEENDENCIES ?= ceph
VERSION_PACKAGE := go-ecm/internal/goecmserver/constand

GO_LDFLAGS += -X $(VERSION_PACKAGE).MainPackageVersion="v6.1.0.0"

GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"

BUILD_TIMESTAMP := $(shell date +'%Y%m%d%H%M%S')

STATIK ?= statik

DOCKER_PACKAGE_DOWNLOAD_URL ?= http://x.y/docker-ce.rpm

COMMA := ,
SPACE :=
SPACE +=

TOOLS ?= $(BLOCKER_TOOLS)
DEPS ?= $(DEPEENDENCIES)
DOWNLOAD_DIR ?= $(ROOT_DIR)/internal/goecmagent/static

GO := go
SHELL := /bin/bash
BINS ?= goecm


PLATFORMS := linux_amd64 linux_arm64

ifeq ($(GOOS), windows)
	$(error unsupported platform. Please build binary from unix)
endif

ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif

	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
endif


.PHONY: build
build: build.server
	@echo "========================> Build server docker image"
	@docker build -t x/goecmserver:$(BUILD_TIMESTAMP) .
	@echo "========================> End of build server docker image"
	echo "========================> push image"
	@docker push x/goecmserver:$(BUILD_TIMESTAMP)
	@echo "========================> Done"


.PHONY: build.server
build.server: download.dependency.docker npm.build gender.template
	@echo "========================> Build server binary from local"
	@mkdir -p $(OUTPUT_DIR)/platforms/linux/amd64/goecmserver
	@$(GO) build $(GO_BUILD_FLAGS) -o goecmserver cmd/goecmserver/goecmserver.go

.PHONY: build.agent
build.agent: download.dependency.docker
	@$(GO) build -o goecmagent  cmd/goecmagent/goecmagent.go
	@$(shell rpmbuild -bb goecmagent.spec)

.PHONY: npm.build
npm.build: tools.verify.statik
	@echo "========================> Run build web resource"
	@rm -fr /root/.npm/*
	@cd web && npm i -g --unsafe-perm && npm i --unsafe-perm && ./node_modules/vite/bin/vite.js build
	@echo "========================> Run package web resource"
	@statik -src=./web/dist

.PHONY: clean
clean:
	@echo "========================>"
	@-rm -vrf $(OUTPUT_DIR)
	@-rm -vrf ./web/node_modules
	@-rm -vrf ./web/dist


.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "========================> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if !which $* &> /dev/null; then $(MAKE) tools.install.$*;fi

.PHONY: install.statik
install.statik:
	@$(GO) install github.com/rakyll/statik@latest

.PHONY: gender.template
gender.template:
	@echo "========================> Starting generate compose template file"
	@$(shell bash generate.sh v6.1.0.0)
	@echo "========================> End of generate compose template file"

.PHONY: download.dependency.%
download.dependency.%:
	@echo "========================> Downloading $*"
	@$(MAKE) download.$*

.PHONY: download.docker
download.docker:
	@echo "========================> Starting download docker from $(DOCKER_PACKAGE_DOWNLOAD_URL)"
	@curl -Lk "$(DOCKER_PACKAGE_DOWNLOAD_URL)" -o $(DOWNLOAD_DIR)/docker.rpm
	@echo "========================> End of download docker"
