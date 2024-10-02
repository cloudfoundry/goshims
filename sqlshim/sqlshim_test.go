package sqlshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/sqlshim"
)

var _ = Describe("Sqlshim", func() {
	It("should implement the interface", func() {
		var obj sqlshim.Sql
		obj = &sqlshim.SqlShim{}
		Expect(obj).To(Equal(&sqlshim.SqlShim{}))
	})

})
