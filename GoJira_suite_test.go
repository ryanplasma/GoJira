package GoJira_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"testing"
)

func TestGoJira(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoJira Suite")
}

func respondWithJSON(w http.ResponseWriter, s string) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	fmt.Fprintf(w, s)
}

func loadFixture(f string) string {
	pwd, _ := os.Getwd()
	p := path.Join(pwd, "fixtures", f)
	c, _ := ioutil.ReadFile(p)
	return string(c)
}
