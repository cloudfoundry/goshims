package usershim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsershim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usershim Suite")
}
