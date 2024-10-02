package bufioshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/bufioshim"
)

var _ = Describe("Bufioshim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj bufioshim.Bufio
		obj = &bufioshim.BufioShim{}
		Expect(obj).To(Equal(&bufioshim.BufioShim{}))
	})

})
