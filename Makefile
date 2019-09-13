# Copyright © 2017 Heptio
# Copyright © 2017 Craig Tracey
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROJECT := gangway
# Where to push the docker image.
REGISTRY ?= gcr.io/heptio-images
IMAGE := $(REGISTRY)/$(PROJECT)
SRCDIRS := ./cmd/gangway
PKGS := $(shell go list ./cmd/... ./internal/...)

VERSION ?= master

all: build

build: deps bindata
	go build ./...

install:
	go install -v ./cmd/gangway/...

setup:
	go get -u github.com/mjibson/esc/...
	curl -o assets/prism-bash.js https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/components/prism-bash.js
	curl -o assets/prism.js https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/prism.js
	curl -o assets/prism.css https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/themes/prism.css
	curl -o assets/prism-powershell.js https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/components/prism-powershell.js
	curl -o assets/materialize.min.css https://raw.githubusercontent.com/Dogfalo/materialize/v1-dev/dist/css/materialize.min.css
	curl -o assets/materialize.min.js https://raw.githubusercontent.com/Dogfalo/materialize/v1-dev/dist/js/materialize.min.js

check: test vet gofmt staticcheck misspell

deps:
	GO111MODULE=on go mod tidy && GO111MODULE=on go mod vendor && GO111MODULE=on go mod verify

vet: | test
	go vet ./...

bindata:
	esc -o cmd/gangway/bindata.go templates/ assets/

test:
	go test -v ./...

staticcheck:
	@go get honnef.co/go/tools/cmd/staticcheck
	staticcheck $(PKGS)

misspell:
	@go get github.com/client9/misspell/cmd/misspell
	misspell \
		-i clas \
		-locale US \
		-error \
		cmd/* docs/* *.md

gofmt:
	@echo Checking code is gofmted
	@test -z "$(shell gofmt -s -l -d -e $(SRCDIRS) | tee /dev/stderr)"

image:
	docker build . -t $(IMAGE):$(VERSION)

push:
	docker push $(IMAGE):$(VERSION)

.PHONY: all deps bindata test image setup
