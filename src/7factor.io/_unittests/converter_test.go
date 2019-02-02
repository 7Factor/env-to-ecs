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

// Constants for testing.
const EmptyEnvironmentArray = `[]`
const SingleLineSingleInput = `A=B`
const MultiLineInput = `
A=B
D=E
`
const SingleLineMultiInput = `W=X Y=Z`

var _ = Describe("The ECS converter", func() {
	Context("When passed a blank file", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.TransformAndTranslate("")
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(EmptyEnvironmentArray))
		})
	})

	Context("When passed a single line file without comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(SingleLineSingleInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"}]`))
		})
	})

	Context("When passed a multi line file with newlines in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(MultiLineInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"},{"name":"D","value":"E"}]`))
		})
	})

	Context("When passed a single line file with multiple items", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(SingleLineMultiInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"W","value":"X"},{"name":"Y","value":"Z"}]`))
		})
	})
})
