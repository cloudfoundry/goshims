package timeshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimeshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timeshim Suite")
}
