package generic_test

import (
	"crypto/tls"
	"fmt"
	"net/http"

	. "tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kubelet", func() {
	It("Should fail when unauthenticated requests are made to kubelet", func() {
		firstWorkerIP, err := GetNodeIP()
		Expect(err).NotTo(HaveOccurred())
		endpoint := fmt.Sprintf("https://%s:10250/pods", firstWorkerIP)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(endpoint)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(401))
	})
})