package _unittests

import (
	"7factor.io/args"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var mockNoArgs = []string{}
var mockInfileArg = []string{"valid_path.env"}
var mockWithOutputShortFlag = []string{"valid_path.env", "-o"}
var mockWithOutputLongFlag = []string{"valid_path.env", "--output"}

var _ = Describe("The argument parser", func() {
	Context("When not passed `INFILE` arg", func() {
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
			Expect(config.InFile).To(Equal("valid_path.env"))
		})
	})

	Context("The `-o, --output` flag", func() {
		It("Errors with empty argument.", func() {
			os.Args = mockWithOutputShortFlag
			_, err := args.GetArguments()
			Expect(err).ToNot(BeNil())

			os.Args = mockWithOutputLongFlag
			_, err = args.GetArguments()
			Expect(err).ToNot(BeNil())
		})
	})
})
