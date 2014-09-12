package GoJira_test

import (
	. "github.com/ryanplasma/GoJira"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Issue", func() {
	var (
		mux     *http.ServeMux
		server  *httptest.Server
		session *Session
	)

	BeforeEach(func() {
		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		session = NewSession(server.URL, "TEST", "TEST", "TEST")
	})

	Describe("Issue()", func() {
		It("should return and Issue", func() {
			mux.HandleFunc("/issue/TEST-123", func(w http.ResponseWriter, r *http.Request) {
				respondWithJSON(w, loadFixture("issue.json"))
			})
			issue, err := session.Issue("TEST-123")
			if err != nil {
				Fail("ERROR: " + err.Error())
			}
			Expect(issue).Should(BeAssignableToTypeOf(&Issue{}))
			Expect(issue.Key).ToNot(BeNil())
		})
	})
})
