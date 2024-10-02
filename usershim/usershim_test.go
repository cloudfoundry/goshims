package usershim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/usershim"
)

var _ = Describe("Usershim", func() {
	It("should implement the interface", func() {
		var obj usershim.User
		obj = &usershim.UserShim{}
		Expect(obj).To(Equal(&usershim.UserShim{}))
	})

})
