package redis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UniversalClient", func() {
	var client UniversalClient

	AfterEach(func() {
		if client != nil {
			Expect(client.Close()).To(Succeed())
		}
	})

	It("should connect to failover servers", func() {
		client = NewUniversalClient(&UniversalOptions{
			MasterName: sentinelName,
			Addrs:      []string{":" + sentinelPort},
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

	It("should connect to simple servers", func() {
		client = NewUniversalClient(&UniversalOptions{
			Addrs: []string{redisAddr},
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

	It("should connect to clusters", func() {
		client = NewUniversalClient(&UniversalOptions{
			Addrs: cluster.addrs(),
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

})
