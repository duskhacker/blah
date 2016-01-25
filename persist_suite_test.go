package persist

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	suite = "Persist"
	dao   dataAccessObject
)

func TestPersist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, suite+" Suite")
}

var _ = BeforeSuite(func() {
	per, err := Open("args")
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	dao = per.(dataAccessObject)
})
