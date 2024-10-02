package execshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/execshim"
)

var _ = Describe("Execshim", func() {
	It("should implement the interface", func() {
		var obj execshim.Exec
		obj = &execshim.ExecShim{}
		Expect(obj).To(Equal(&execshim.ExecShim{}))
	})

})
