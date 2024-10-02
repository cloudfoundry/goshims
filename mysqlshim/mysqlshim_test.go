package mysqlshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/mysqlshim"
)

var _ = Describe("Mysqlshim", func() {
	It("should implement the interface", func() {
		var obj mysqlshim.MySQL
		obj = &mysqlshim.MySQLShim{}
		Expect(obj).To(Equal(&mysqlshim.MySQLShim{}))
	})

})
