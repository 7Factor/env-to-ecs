package _unittests

import (
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The file reader/converter", func() {
	Context("When passed an invalid file path", func() {
		It("Tells us if the file doesn't exist.", func() {
			_, err := converter.ReadAndConvert("/path/does/not/exist/ever")
			Expect(err).ToNot(BeNil())
		})
	})

	Context("When passed a valid file path", func() {
		It("Returns the contents of the path and no error", func() {
			contents, err := converter.ReadAndConvert("valid_path.env")
			Expect(err).To(BeNil())
			Expect(contents).ToNot(BeEmpty())
		})
	})
})

const EmptyEnvironmentArray = `[]`
var _ = Describe("The ECS converter", func() {
	Context("When passed a blank file", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.Transform("")
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(EmptyEnvironmentArray))
		})
	})

	Context("When passed a single line file without comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.Transform("A=B")
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"}]`))
		})
	})
})