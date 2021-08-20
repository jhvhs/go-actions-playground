GOBIN=$(shell pwd)/bin
GOPATH=$(shell pwd)/vendor:$(shell pwd)

all: test

install_prep:
	go mod vendor

install_ginkgo: install_prep
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get github.com/onsi/ginkgo/ginkgo

test: install_ginkgo
	$(GOBIN)/ginkgo -r .
