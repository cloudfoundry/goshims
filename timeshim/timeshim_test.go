package timeshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/timeshim"
)

var _ = Describe("Timeshim", func() {
	It("should implement the interface", func() {
		//lint:ignore S1021 The purpose of this test is to keep it seperate
		var obj timeshim.Time
		obj = &timeshim.TimeShim{}
		Expect(obj).To(Equal(&timeshim.TimeShim{}))
	})

})
