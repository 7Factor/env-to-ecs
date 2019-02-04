package _unittests

import (
	"7factor.io/args"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var mockNoArgs = []string{"cmd"}
var mockInfileArg = []string{"cmd", "valid_path.env"}

var _ = Describe("The argument parser", func() {
	Context("When passed no arguments", func() {
		It("Returns an error.", func() {
			os.Args = mockNoArgs
			_, err := args.GetArguments()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("When passed an infile and nothing else", func() {
		It("Returns a blank config struct with the input file filed out.", func() {
			os.Args = mockInfileArg
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.EnvironmentFile).To(Equal("valid_path.env"))
		})
	})
})
