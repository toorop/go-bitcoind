package bitcoind

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBitcoind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bitcoind Suite")
}
