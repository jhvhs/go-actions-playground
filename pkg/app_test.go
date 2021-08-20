package pkg_test

import (
	"github.com/jhvhs/go-actions-playground/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("App", func() {
	Expect(true)

	var (
		stdout  *gbytes.Buffer
		subject pkg.Application
	)

	BeforeEach(func() {
		stdout = &gbytes.Buffer{}
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
