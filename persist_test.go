package persist

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Persist", func() {
	Describe("ServerInfo", func() {
		It("returns server information", func() {
			info, err := dao.ServerInfo()
			Expect(err).ToNot(HaveOccurred())
			Expect(info).To(Equal("info"))
		})
	})
})
