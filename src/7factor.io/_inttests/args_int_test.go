package _inttests

import (
	"7factor.io/args"
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var mockInfileArg = []string{"cmd", "valid_path.env"}
var mockWithOutputShortFlag = []string{"cmd", "valid_path.env", "-o", "output.json"}

var _ = Describe("Calling converter with args", func() {
	Context("When passing args to the Converter", func() {
		It("Passes the inFile with no errors", func() {
			os.Args = mockInfileArg
			args, err := args.GetArguments()
			Expect(err).To(BeNil())
			_, err = converter.ReadAndConvert(args.InFile, "")
			Expect(err).To(BeNil())
		})

		It("Passes the outFile with no errors", func() {
			os.Args = mockWithOutputShortFlag
			args, err := args.GetArguments()
			Expect(err).To(BeNil())
			outfile, err := converter.ReadAndConvert(args.InFile, args.OutFile)
			Expect(err).To(BeNil())
			Expect(outfile).To(Equal("output.json"))
		})
	})
})
