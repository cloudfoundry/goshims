package syscallshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/syscallshim"
)

var _ = Describe("Syscallshim", func() {
	It("should implement the interface", func() {
		var obj syscallshim.Syscall
		obj = &syscallshim.SyscallShim{}
		Expect(obj).To(Equal(&syscallshim.SyscallShim{}))
	})

})
