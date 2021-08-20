all: test

install_ginkgo:
	@go get github.com/onsi/ginkgo

test: install_ginkgo
	@ginkgo -r .
