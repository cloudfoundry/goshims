package ldapshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLdapshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ldapshim Suite")
}
