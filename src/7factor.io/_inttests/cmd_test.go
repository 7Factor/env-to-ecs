package _inttests

import (
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

var _ = Describe("Compiling and running the script with arguments", func() {
	BeforeSuite(func() {
		var err error
		pathToCMD, err = Build("7factor.io/cmd")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("When script is called with no INFILE", func() {
		It("Errors and exits in the expected manner.", func() {
			command := exec.Command(pathToCMD)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(1))
			Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
		})
	})

	Context("When the script is called with an INFILE that does no exist", func() {
		It("Errors and exits in the expected manner.", func() {
			command := exec.Command(pathToCMD, "-i", "/path/does/not/exist/ever")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(1))
			Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
		})
	})

	Context("When calling the script with INFILE only", func() {
		It("Prints the output to stdout", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Expect(session).Should(Exit(0))
			Eventually(session.Err.Contents()).Should(BeEmpty())
			Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutput))
		})
	})

	Context("When calling the script with --output but no specified outfile", func() {
		It("Errors and exits in the expected manner.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(1))
			Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
		})
	})

	Context("When calling the script with --output and passing a valid output file", func() {
		It("Writes to the outfile with no errors and exits cleanly.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o", "output.json")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(0))
			Eventually(session.Err.Contents()).Should(BeEmpty())

			actual, err := ioutil.ReadFile("output.json")
			Expect(err).To(BeNil())
			Expect(string(actual)).To(Equal(expectedOutput))
		})
	})

	Context("When calling the script with -v and passing a valid variable", func() {
		It("Writes the correct output with no errors and exits cleanly.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-v", "extra_var=a_database_connection_string")

			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())

			session.Wait()

			Eventually(session).Should(Exit(0))
			Eventually(session.Err.Contents()).Should(BeEmpty())

			Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithExtraVar))
		})
	})

	Context("When calling the script with multiple -v flags and passing valid variables", func() {
		It("Writes the correct output with no errors and exits cleanly.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-v", "extra_var1=this_thing", "-v", "extra_var2=that_thing")

			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())

			session.Wait()

			Eventually(session).Should(Exit(0))
			Eventually(session.Err.Contents()).Should(BeEmpty())

			Expect(string(session.Out.Contents())).To(ContainSubstring(expectedOutputWithMultipleExtraVar))
		})
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})
})
