package filepathshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/filepathshim"
)

var _ = Describe("Filepathshim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj filepathshim.Filepath
		obj = &filepathshim.FilepathShim{}
		Expect(obj).To(Equal(&filepathshim.FilepathShim{}))
	})

})
