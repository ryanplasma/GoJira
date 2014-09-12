package GoJira_test

import (
	. "github.com/ryanplasma/GoJira"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoJira", func() {
	Describe("NewSession()", func() {
		It("should return a session", func() {
			session := NewSession("TEST", "TEST", "TEST", "")
			Expect(session).Should(BeAssignableToTypeOf(&Session{}))
		})
	})
})
