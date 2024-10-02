package osshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/osshim"
)

var _ = Describe("Osshim", func() {
	It("should implement the file interface", func() {
		var obj osshim.File
		obj = &osshim.FileShim{}
		Expect(obj).To(Equal(&osshim.FileShim{}))
	})
	It("should implement the os interface", func() {
		var obj osshim.Os
		obj = &osshim.OsShim{}
		Expect(obj).To(Equal(&osshim.OsShim{}))
	})

})
