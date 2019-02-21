package _unittests

import (
	"7factor.io/args"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var mockNoArgs = []string{"cmd"}
var mockInfileArg = []string{"cmd", "-i", "valid_path.env"}
var errorMockWithOutputShortFlag = []string{"cmd","valid_path.env", "-o"}
var errorMockWithOutputLongFlag = []string{"cmd", "valid_path.env", "--output"}
var mockWithOutputShortFlag = []string{"cmd", "valid_path.env", "-o", "output.json"}
var mockWithOutputLongFlag = []string{"cmd", "valid_path.env", "--output", "output.json"}

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
	//
	//Context("When called with output flag and no specified output file", func() {
	//	It("Errors in the expected manner", func() {
	//		os.Args = errorMockWithOutputShortFlag
	//		_, err := args.GetArguments()
	//		Expect(err).ToNot(BeNil())
	//
	//		os.Args = errorMockWithOutputLongFlag
	//		_, err = args.GetArguments()
	//		Expect(err).ToNot(BeNil())
	//	})
	//
	//	It("Returns the config struct with the output file filled out.", func() {
	//		os.Args = mockWithOutputShortFlag
	//		config, err := args.GetArguments()
	//		Expect(err).To(BeNil())
	//		Expect(config.OutFile).To(Equal("output.json"))
	//
	//		os.Args = mockWithOutputLongFlag
	//		config, err = args.GetArguments()
	//		Expect(err).To(BeNil())
	//		Expect(config.OutFile).To(Equal("output.json"))
	//	})
	//})
})
