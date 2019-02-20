package _unittests

import (
	"7factor.io/args"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var mockNoArgs = []string{"cmd"}
var mockInfileArg = []string{"cmd", "-i", "valid_path.env"}
var mockWithOutput = []string{"cmd", "-i", "valid_path.env", "-o", "output.json"}

var _ = Describe("The argument parser", func() {
	Context("When not passed `INFILE` arg", func() {
		It("Returns an error.", func() {
			os.Args = mockNoArgs
			_, err := args.GetArguments()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("When passed an infile and nothing else", func() {
		It("Returns the infile and the the outfile as `stdout`.", func() {
			os.Args = mockInfileArg
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.InFile).To(Equal("valid_path.env"))
		})

		It("Prints the output to stdout.", func() {
			os.Args = mockInfileArg
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.OutFile).To(Equal("stdout"))
		})
	})

	Context("When called with the output flag and a specified output file", func() {
		It("Returns appropriate `ArgConfig` struct with the correct outfile.", func() {
			os.Args = mockWithOutput
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.OutFile).To(Equal("output.json"))
		})
	})

	Context("When called with the additional values flag and no specified additional value", func() {
		It("errors in the expected manner", func() {
			os.Args = errorMockWithAdditionalValuesShortFlag
			_, err := args.GetArguments()
			Expect(err).ToNot(BeNil())

			os.Args = errorMockWithAdditionalValuesLongFlag
			_, err = args.GetArguments()
			Expect(err).ToNot(BeNil())
		})
	})
})
