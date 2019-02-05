package _inttests

import (
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The cmd script", func() {
	Context("When the script is called it writes the correct output to the outfile", func() {
		It("Creates the outfile if it does not exist", func() {
			_, err := converter.ReadAndConvert("valid_path.env", "output.json")
			Expect(err).To(BeNil())
			Expect("output.json").Should(BeAnExistingFile())
		})
	})
})
