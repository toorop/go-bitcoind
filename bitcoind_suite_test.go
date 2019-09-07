package navcoind

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNavcoind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Navcoind Suite")
}
