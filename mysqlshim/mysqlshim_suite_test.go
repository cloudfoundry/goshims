package mysqlshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMysqlshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysqlshim Suite")
}
