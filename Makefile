TRAG.Gopkg:=earth
TRAG.Version:=$(TRAG.Gopkg)/pkg/version

RUN_IN_DOCKER:=docker run -it -v `pwd`:/go/src/$(TRAG.Gopkg) -v `pwd`/tmp/mod:/go/pkg/mod -v `pwd`/tmp/cache:/root/.cache/go-build  -w /go/src/$(TRAG.Gopkg) -e GOPROXY=https://goproxy.io -e GOBIN=/go/src/$(TRAG.Gopkg)/tmp/bin -e USER_ID=`id -u` -e GROUP_ID=`id -g` $(BUILDER_IMAGE)
RUN_IN_DOCKER_WITHOUT_GOPATH:=docker run -it -v `pwd`:/go/src/$(TRAG.Gopkg) -v `pwd`/tmp/cache:/root/.cache/go-build  -w /go/src/$(TRAG.Gopkg) -e USER_ID=`id -u` -e GROUP_ID=`id -g` $(BUILDER_IMAGE)
GO_BUILD_DARWIN:=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags netgo
GO_BUILD_LINUX:=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo
GO_BUILD_WINDOWS:=CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags netgo
GO_FMT:=goimports -l -w -e -local=earth -srcdir=/go/src/$(TRAG.Gopkg)
GO_RACE:=go build -race
GO_VET:=go vet

define get_diff_files
    $(eval DIFF_FILES=$(shell git diff --name-only --diff-filter=ad | grep -E "^(test|cmd|pkg)/.+\.go"))
endef
define get_build_flags
    $(eval SHORT_VERSION=$(shell git describe --tags --always --dirty="-dev"))
    $(eval SHA1_VERSION=$(shell git show --quiet --pretty=format:%H))
	$(eval DATE=$(shell date +'%Y-%m-%dT%H:%M:%S'))
	$(eval BUILD_FLAG= -X $(TRAG.Version).ShortVersion="$(SHORT_VERSION)" \
		-X $(TRAG.Version).GitSha1Version="$(SHA1_VERSION)" \
		-X $(TRAG.Version).BuildDate="$(DATE)" \
		-w -s)
endef


.PHONY: generate-in-local
generate-in-local: ## Generate code from protobuf file in local
	cd ./apis && make generate


.PHONY: compose-up
compose-up:
	docker-compose up -d --remove-orphans

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans