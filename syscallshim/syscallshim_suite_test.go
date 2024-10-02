package syscallshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSyscallshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Syscallshim Suite")
}
