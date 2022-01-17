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
var mockWithVariable = []string{"cmd", "-i", "valid_path.env", "-v", "A=B"}
var mockWithMultipleVariables = []string{"cmd", "-i", "valid_path.env", "-v", "C=D", "-v", "E=F"}
var mockWithJsonParent = []string{"cmd", "-i", "valid_path.env", "-p", "env"}

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

		It("Does not set the json parent key.", func() {
			os.Args = mockInfileArg
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.JsonParent).To(BeNil())
		})
	})

	Context("When called with the output flag and a specified output file", func() {
		It("Returns the appropriate `ArgConfig` struct with the correct outfile.", func() {
			os.Args = mockWithOutput
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.OutFile).To(Equal("output.json"))
		})
	})

	Context("When called with the variable flag and a specified variable", func() {
		It("Returns then appropriate `ArgConfig` struct with the expected variable", func() {
			os.Args = mockWithVariable
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.Variables).To(Equal([]string{"A=B"}))
		})
	})

	Context("When called with multiple variable flags and specified variables", func() {
		It("Returns the appropriate `ArgConfig` struct with the expected variables", func() {
			os.Args = mockWithMultipleVariables
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(config.Variables).To(Equal([]string{"C=D", "E=F"}))
		})
	})

	Context("When called with an infile and a json parent key", func() {
		It("Sets the json parent key.", func() {
			os.Args = mockWithJsonParent
			config, err := args.GetArguments()
			Expect(err).To(BeNil())
			Expect(*config.JsonParent).To(Equal("env"))
		})
	})
})
