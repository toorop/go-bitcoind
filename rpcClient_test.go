package bitcoind

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

var _ = Describe("RpcClient", func() {
	Describe("Initialise a new rpcClient", func() {
		Context("when initialisation succeeded", func() {
			client, err := newClient("127.0.0.1", 8334, "user", "paswd", false, 30)
			It("err should be nil", func() {
				Expect(err).To(BeNil())
			})
			It("return must be a new rpcClient address ", func() {
				Expect(client).Should(BeAssignableToTypeOf(&rpcClient{}))
			})
		})

		Context("when initialisation failed (empty host)", func() {
			client, err := New("", 8334, "user", "paswd", false)
			It("err should occured", func() {
				Expect(err).Should(HaveOccurred())
			})
			It("rpcClient should be nil", func() {
				Expect(client).To(BeNil())
			})
		})
	})

	Describe("Do requests", func() {
		Context("When connexion fail", func() {
			client, err := newClient("127.0.0.1", 123, "fake", "fake", false, 30)
			_, err = client.call("getdifficulty", nil)
			It("err should occured", func() {
				Expect(err).Should(MatchError("Post http://127.0.0.1:123: dial tcp 127.0.0.1:123: connection refused"))
			})
		})

		Context("When timeout occured", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep((RPCCLIENT_TIMEOUT + 2) * time.Second)
				fmt.Fprintln(w, "Hello, client")
			}))
			defer ts.Close()
			p := strings.Split(ts.URL, ":")
			host := p[1][2:]
			port, err := strconv.ParseInt(p[2], 10, 64)
			client, err := newClient(host, int(port), "fake", "fake", false, 30)
			_, err = client.call("getdifficulty", nil)

			It("timeout err should occured", func() {
				Expect(err).Should(MatchError("Timeout reading data from server"))
			})

		})

	})

})
