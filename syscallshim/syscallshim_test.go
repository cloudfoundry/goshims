package syscallshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/syscallshim"
)

var _ = Describe("Syscallshim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj syscallshim.Syscall
		obj = &syscallshim.SyscallShim{}
		Expect(obj).To(Equal(&syscallshim.SyscallShim{}))
	})

})
