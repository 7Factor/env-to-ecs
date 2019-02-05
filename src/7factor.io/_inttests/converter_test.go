package _inttests

import (
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The file reader/converter", func() {
	Context("When passed an invalid file path", func() {
		It("Tells us if the file doesn't exist.", func() {
			_, err := converter.ReadAndConvert("/path/does/not/exist/ever", "")
			Expect(err).ToNot(BeNil())
		})
	})

	Context("When passed an valid file path", func() {
		It("Creates the outfile if it does not exist", func() {
			_, err := converter.ReadAndConvert("valid_path.env", "output.json")
			Expect(err).To(BeNil())
			Expect("output.json").Should(BeAnExistingFile())
		})

		It("Correctly writes transformed contents to outfile", func() {
			transformedContents, err := converter.ReadAndConvert("valid_path.env", "output.json")
			Expect(err).To(BeNil())
			Expect("output.json").To(ContainSubstring(transformedContents))
		})
	})
})
