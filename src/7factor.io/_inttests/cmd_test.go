package _inttests

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"io/ioutil"
	"os/exec"
)

var pathToCMD = "7factor.io/cmd"

var expectedOutput = `[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"}]`
var expectedOutputWithExtraVar = `[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"},{"name":"extra_var","value":"a_database_connection_string"}]`
var expectedOutputWithMultipleExtraVar = `[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"},{"name":"extra_var1","value":"this_thing"},{"name":"extra_var2","value":"that_thing"}]`

var jsonParentName = `env_vars`
var expectedOutputWithJsonParent = fmt.Sprintf(`{"%s":%s}`, jsonParentName, expectedOutput)
var expectedOutputWithJsonParentNoVars = fmt.Sprintf(`{"%s":[]}`, jsonParentName)

var _ = Describe("Compiling and running the script with arguments", func() {
	BeforeSuite(func() {
		var err error
		pathToCMD, err = Build("7factor.io/cmd")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("The -i | --infile flag", func() {
		Context("When script is called with no INFILE", func() {
			It("Errors and exits in the expected manner.", func() {
				command := exec.Command(pathToCMD)
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxErrorCode))
				Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
			})
		})

		Context("When the script is called with an INFILE that does no exist", func() {
			It("Errors and exits in the expected manner.", func() {
				command := exec.Command(pathToCMD, "-i", "/path/does/not/exist/ever")
				session, err := Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
				session.Wait()

				Eventually(session).Should(Exit(withLinuxErrorCode))
				Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
			})
		})

		Context("When calling the script with INFILE only", func() {
			It("Prints the output to stdout", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env")
				session := setUpSessionAndWait(command)

				Expect(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutput))
			})
		})
	})

	Describe("The -o | --outfile flag", func() {
		Context("When calling the script with --output but no specified outfile", func() {
			It("Errors and exits in the expected manner.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o")
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxErrorCode))
				Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
			})
		})

		Context("When calling the script with --output and passing a valid output file", func() {
			It("Writes to the outfile with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o", "output.json")
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				actual, err := ioutil.ReadFile("output.json")
				Expect(err).To(BeNil())
				Expect(string(actual)).To(Equal(expectedOutput))
			})
		})
	})

	Describe("The -v | --variable flag", func() {
		Context("When calling the script with -v and passing a valid variable", func() {
			It("Writes the correct output with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-v", "extra_var=a_database_connection_string")
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithExtraVar))
			})
		})

		Context("When calling the script with multiple -v flags and passing valid variables", func() {
			It("Writes the correct output with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-v", "extra_var1=this_thing", "-v", "extra_var2=that_thing")
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithMultipleExtraVar))
			})
		})
	})

	Describe("Handling files with no newlines", func() {
		Context("When the script is called with an INFILE that has no newline at the end", func() {
			It("Writes to the outfile correctly with no errors", func() {
				command := exec.Command(pathToCMD, "-i", "with_no_new_line.env")
				session := setUpSessionAndWait(command)

				Expect(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutput))
			})
		})

		Context("When the script is called with an INFILE that has no newline at the end and passing extra variables", func() {
			It("Writes to the outfile correctly with no errors", func() {
				command := exec.Command(pathToCMD, "-i", "with_no_new_line.env", "-v", "extra_var=a_database_connection_string")
				session := setUpSessionAndWait(command)

				Expect(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithExtraVar))
			})
		})
	})

	Describe("The -p | --parent flag", func() {
		Context("When calling the script with parent param", func() {
			It("Writes the correct output with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-p", jsonParentName)
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithJsonParent))
			})
		})

		Context("When calling the script with parent param and empty env file", func() {
			It("Writes the correct output with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "empty.env", "-p", jsonParentName)
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxPassingCode))
				Eventually(session.Err.Contents()).Should(BeEmpty())

				Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithJsonParentNoVars))
			})
		})

		Context("When calling the script with parent param with blank arg", func() {
			It("Writes the correct output with no errors and exits cleanly.", func() {
				command := exec.Command(pathToCMD, "-i", "valid_path.env", "-p", "\"\"")
				session := setUpSessionAndWait(command)

				Eventually(session).Should(Exit(withLinuxErrorCode))
				Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
			})
		})
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})
})
