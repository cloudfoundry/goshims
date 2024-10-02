package usershim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/usershim"
)

var _ = Describe("Usershim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj usershim.User
		obj = &usershim.UserShim{}
		Expect(obj).To(Equal(&usershim.UserShim{}))
	})

})
