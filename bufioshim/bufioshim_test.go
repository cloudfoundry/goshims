package bufioshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/bufioshim"
)

var _ = Describe("Bufioshim", func() {
	It("should implement the interface", func() {
		var obj bufioshim.Bufio
		obj = &bufioshim.BufioShim{}
		Expect(obj).To(Equal(&bufioshim.BufioShim{}))
	})

})
