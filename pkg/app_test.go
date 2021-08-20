package pkg_test

import (
	"io"

	"github.com/jhvhs/go-actions-playground/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("App", func() {
	Expect(true)

	var (
		stdout  io.Writer
		subject pkg.Application
	)

	BeforeEach(func() {
		stdout = gbytes.NewBuffer()
		subject = pkg.Application{
			Name:   "John",
			Output: stdout,
		}
	})

	It("should produce the output", func() {
		subject.Run()

		Expect(stdout).To(gbytes.Say("Hello, John"))
	})
})
