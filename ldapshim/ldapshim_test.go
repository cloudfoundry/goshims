package ldapshim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/goshims/ldapshim"
)

var _ = Describe("Ldapshim", func() {
	It("should implement the interface", func() {
		var obj ldapshim.Ldap
		obj = &ldapshim.LdapShim{}
		Expect(obj).To(Equal(&ldapshim.LdapShim{}))
	})

})
