package _unittests

import (
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Constants for testing.
const EmptyEnvironmentArray = `[]`
const SingleLineSingleInput = `A=B`
const MultiLineInput = `
A=B
D=E
`
const MultiLineWithSpacesInput = `

L=M

N=O
`
const SingleLineMultiInput = `W=X Y=Z`
const illegalJSONInput = `A=\"`
const MultiLineWithComments = `
# this is a comment
Q=R
S=T
`

const InputWithHash = `WITHHASH=#FOO#`

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

	Context("When passed a multi-line file with newlines in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(MultiLineInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"},{"name":"D","value":"E"}]`))
		})
	})

	Context("When passed a multi-line file with spaces in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(MultiLineWithSpacesInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"L","value":"M"},{"name":"N","value":"O"}]`))
		})
	})

	Context("When passed a single line file with multiple items", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(SingleLineMultiInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"W","value":"X"},{"name":"Y","value":"Z"}]`))
		})
	})

	Context("When passed characters that cause the JSON translation to fail", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(illegalJSONInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"\\\""}]`))
		})
	})

	Context("When passed a multi-line file with comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(MultiLineWithComments)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"Q","value":"R"},{"name":"S","value":"T"}]`))
		})
	})

	Context("When passed a file with `#` in the `name` and `value` params and is otherwise valid", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.TransformAndTranslate(InputWithHash)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"WITHHASH","value":"#FOO#"}]`))
		})
	})
})
