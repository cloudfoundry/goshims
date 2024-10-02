package sqlshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSqlshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sqlshim Suite")
}
