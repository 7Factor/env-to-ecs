package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"os/exec"
	"testing"
)

const (
	withLinuxPassingCode = 0
	withLinuxErrorCode   = 1
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ENV to ECS Integration Test Suite")
}

func setUpSessionAndWait(command *exec.Cmd) *Session {
	session, err := Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ShouldNot(HaveOccurred())
	session.Wait()
	return session
}
