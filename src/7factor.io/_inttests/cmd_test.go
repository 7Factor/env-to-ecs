package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"io/ioutil"
	"os/exec"
)

var pathToCMD = "7factor.io/cmd"

var _ = Describe("Compiling and running the script with arguments", func() {
	BeforeSuite(func() {
		var err error
		pathToCMD, err = Build("7factor.io/cmd")
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
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
			Eventually(session.Out).Should(gbytes.Say(`[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"}]`))
		})
	})

	Context("When calling the script with -o but no specified outfile", func() {
		It("Errors and exits in the expected manner.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(1))
			Eventually(session.Err.Contents()).ShouldNot(BeEmpty())
		})
	})

	Context("When calling the script with -o and passing a valid output file", func() {
		It("Writes to the outfile with no errors and exits cleanly.", func() {
			command := exec.Command(pathToCMD, "-i", "valid_path.env", "-o", "output.json")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			session.Wait()
			Eventually(session).Should(Exit(0))
			Eventually(session.Err.Contents()).Should(BeEmpty())

			actual, err := ioutil.ReadFile("output.json")
			Expect(err).To(BeNil())
			Expect(string(actual)).To(Equal(`[{"name":"FOO","value":"bar"},{"name":"BAZ","value":"boo"}]`))
		})
	})
})
