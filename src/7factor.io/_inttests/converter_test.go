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

	Context("When passed a valid file path", func() {
		It("Returns the contents of the path and no error", func() {
			contents, err := converter.ReadAndConvert("valid_path.env", "")
			Expect(err).To(BeNil())
			Expect(contents).ToNot(BeEmpty())
		})
	})

	Context("When attempting the ECS converter", func() {
		It("Returns converted contents in the expected manner", func() {
			contents, err := converter.ReadAndConvert("valid_path.env", "")
			Expect(err).To(BeNil())
			Expect(contents).To(Equal(`[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"}]`))
		})
	})
})
