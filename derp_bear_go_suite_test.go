package derp_bear_go_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDerpBearGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DerpBearGo Suite")
}
