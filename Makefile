PACKAGES=$(shell go list ./... | grep -v '/vendor/')
BUILD_TAGS?=noah
BUILD_FLAGS=-ldflags "-s -w -X noah/version.GitCommit=`git rev-parse --short=8 HEAD`" -mod=vendor

all: build test install

########################################
### Build

build:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o build/noah ./cmd/noah/

build_c:
	CGO_ENABLED=1 go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS) gcc cleveldb' -o build/noah ./cmd/noah/

install:
	CGO_ENABLED=0 go install $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' ./cmd/noah


########################################
### Tools & dependencies

test:
	CGO_ENABLED=1 CGO_LDFLAGS="-lsnappy" go test --count 1 --tags "gcc cleveldb" ./...

create_vendor:
	@echo "--> Installing vendor"
	CGO_ENABLED=0 go mod vendor

########################################
### Formatting, linting, and vetting

fmt:
	@go fmt ./...

metalinter:
	@echo "--> Running linter"
	@gometalinter.v2 --vendor --deadline=600s --disable-all  \
		--enable=deadcode \
		--enable=gosimple \
	 	--enable=misspell \
		--enable=safesql \
		./...
		#--enable=gas \
		#--enable=maligned \
		#--enable=dupl \
		#--enable=errcheck \
		#--enable=goconst \
		#--enable=gocyclo \
		#--enable=goimports \
		#--enable=golint \ <== comments on anything exported
		#--enable=gotype \
	 	#--enable=ineffassign \
	   	#--enable=interfacer \
	   	#--enable=megacheck \
	   	#--enable=staticcheck \
	   	#--enable=structcheck \
	   	#--enable=unconvert \
	   	#--enable=unparam \
		#--enable=unused \
	   	#--enable=varcheck \
		#--enable=vet \
		#--enable=vetshadow \

metalinter_all:
	@echo "--> Running linter (all)"
	gometalinter.v2 --vendor --deadline=600s --enable-all --disable=lll ./...

###########################################################
### Docker image

build-docker:
	cp build/noah DOCKER/noah
	cd DOCKER && make build
	rm -f noah

push-docker:
	cd DOCKER && make push

###########################################################
### Local testnet using docker

# Build linux binary on other platforms
build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

build-compress:
	upx build/noah