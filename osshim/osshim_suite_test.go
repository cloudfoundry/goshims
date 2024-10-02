package osshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOsshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Osshim Suite")
}
