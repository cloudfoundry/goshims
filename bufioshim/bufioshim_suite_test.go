package bufioshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBufioshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bufioshim Suite")
}
