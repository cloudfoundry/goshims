package sqlshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/sqlshim"
)

var _ = Describe("Sqlshim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj sqlshim.Sql
		obj = &sqlshim.SqlShim{}
		Expect(obj).To(Equal(&sqlshim.SqlShim{}))
	})

})
