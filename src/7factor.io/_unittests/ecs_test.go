package _unittests

import (
	"7factor.io/converter"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Constants for testing.
const EmptyEnvironmentArray = `[]`
const EmptyEnvironmentObject = `{}`
const DefaultParentKey = `env_vars`
const EmptyEnvironmentArrayParentKey = `{"%s":[]}`
const SingleLineSingleInput = `A=B`
const EmptyValueInput = `A=`
const MultiEmptyValueInput = `
A=
B=
`
const MultiLineInput = `
A=B
D=E
`
const MultiLineWithSpacesInput = `

L=M

N=O
`
const SingleLineMultiInput = `W=X   Y=Z`
const InputWithSpaces = `X = 1 Y = 1`
const illegalJSONInput = `A=\"`
const MultiLineWithComments = `
# this is a comment
Q=R
S=T
`

// Base64 strings always end with an equals, handle this case.
const InputWithEquals = `A=abcdefg=`
const InputWithMultipleEquals = `A=abcdefg===`
const InputWithQueryParams = `SPRING_DATASOURCE_URL=jdbc:mysql://example.com/database?useSSL=false`
const InputWithQuotedValueQueryParams = `SPRING_DATASOURCE_URL="jdbc:mysql://example.com/database?useSSL=false"`
const MultiLineInputWithEquals = `
A=abcdefg=
B=C
`
const InputWithHash = `WITHHASH=#FOO#`

const InputWithQuotes = `WITHQUOTES="this is a test"`
const InputWithQuotesAndSpaces = `A = "test string"`
const MultiLineInputWithNewlinesInQuotes = `A="test
string"`
const MultiLineInputWithNewlinesInSingleQuotes = `A='test
string'`

const CrazyInput = `A=1 B = 2 C="test string" D = "another test string" E="1" F = "2"
G = "another test string"
H="test string"`

var _ = Describe("The ECS converter", func() {
	Context("When passed a blank file", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{})
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(EmptyEnvironmentArray))
		})
	})

	Context("When passed a single line file without comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{SingleLineSingleInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"}]`))
		})
	})

	Context("When passed an empty value input", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{EmptyValueInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":""}]`))
		})
	})

	Context("When passed multiple empty value inputs", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiEmptyValueInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":""},{"name":"B","value":""}]`))
		})
	})

	Context("When passed a single line file with multiple values and spaces between the equals", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithSpaces})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"X","value":"1"},{"name":"Y","value":"1"}]`))
		})
	})

	Context("When passed a single line file with quotes and spaces", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithQuotesAndSpaces})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"test string"}]`))
		})
	})

	Context("When passed a multi-line file with newlines inside quotes", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineInputWithNewlinesInQuotes})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"test\nstring"}]`))
		})
	})

	Context("When passed a multi-line file with newlines inside single quotes", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineInputWithNewlinesInSingleQuotes})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"test\nstring"}]`))
		})
	})

	Context("When passed a multi-line file with newlines in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"},{"name":"D","value":"E"}]`))
		})
	})

	Context("When passed a multi-line file with spaces in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineWithSpacesInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"L","value":"M"},{"name":"N","value":"O"}]`))
		})
	})

	Context("When passed a single line file with multiple items", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{SingleLineMultiInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"W","value":"X"},{"name":"Y","value":"Z"}]`))
		})
	})

	Context("When passed characters that cause the JSON translation to fail", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{illegalJSONInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"\\\""}]`))
		})
	})

	Context("When passed a multi-line file with comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineWithComments})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"Q","value":"R"},{"name":"S","value":"T"}]`))
		})
	})

	Context("When passed a file with `#` in the `name` and `value` params and is otherwise valid", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithHash})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"WITHHASH","value":"#FOO#"}]`))
		})
	})

	Context("When called with a string variable ending with an equals sign", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithEquals})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"abcdefg="}]`))
		})
	})

	Context("When called with a string variable ending with multiple equals signs", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithMultipleEquals})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"abcdefg==="}]`))
		})
	})

	Context("When called with a string variable that is a URL with query params", func() {
		It("Works as expected and doesn't drop query param values.", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithQueryParams})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"SPRING_DATASOURCE_URL","value":"jdbc:mysql://example.com/database?useSSL=false"}]`))
		})
	})

	Context("When called with a string variable that is a quoted URL with query params", func() {
		It("Works as expected and doesn't drop query param values.", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithQuotedValueQueryParams})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"SPRING_DATASOURCE_URL","value":"jdbc:mysql://example.com/database?useSSL=false"}]`))
		})
	})

	Context("When called with a string variable with an equals sign in it and multiple variables", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson([]string{MultiLineInputWithEquals})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"abcdefg="},{"name":"B","value":"C"}]`))
		})
	})

	Context("When called with a quoted string variable", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson([]string{InputWithQuotes})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"WITHQUOTES","value":"this is a test"}]`))
		})
	})

	Context("When called with a string with multiple challenges", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson([]string{CrazyInput})
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"1"},{"name":"B","value":"2"},{"name":"C","value":"test string"},{"name":"D","value":"another test string"},{"name":"E","value":"1"},{"name":"F","value":"2"},{"name":"G","value":"another test string"},{"name":"H","value":"test string"}]`))
		})
	})

	Context("When passed a parent key and a blank file", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.ConvertInputToJsonObject([]string{}, DefaultParentKey)
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(fmt.Sprintf(EmptyEnvironmentArrayParentKey, DefaultParentKey)))
		})
	})

	Context("When passed a single line file without comments and a blank parent key", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.ConvertInputToJsonObject([]string{}, "")
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(EmptyEnvironmentObject))
		})
	})

	Context("When passed a parent key and a single line file without comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJsonObject([]string{SingleLineSingleInput}, DefaultParentKey)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`{"env_vars":[{"name":"A","value":"B"}]}`))
		})
	})
})
