package redis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pipelining", func() {
	var client *Client
	var pipe *Pipeline

	BeforeEach(func() {
		client = NewClient(redisOptions())
		Expect(client.FlushDB().Err()).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(client.Close()).NotTo(HaveOccurred())
	})

	It("supports block style", func() {
		var get *StringCmd
		cmds, err := client.Pipelined(func(pipe Pipeliner) error {
			get = pipe.Get("foo")
			return nil
		})
		Expect(err).To(Equal(Nil))
		Expect(cmds).To(HaveLen(1))
		Expect(cmds[0]).To(Equal(get))
		Expect(get.Err()).To(Equal(Nil))
		Expect(get.Val()).To(Equal(""))
	})

	assertPipeline := func() {
		It("returns no errors when there are no commands", func() {
			_, err := pipe.Exec()
			Expect(err).NotTo(HaveOccurred())
		})

		It("discards queued commands", func() {
			pipe.Get("key")
			pipe.Discard()
			cmds, err := pipe.Exec()
			Expect(err).NotTo(HaveOccurred())
			Expect(cmds).To(BeNil())
		})

		It("handles val/err", func() {
			err := client.Set("key", "value", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			get := pipe.Get("key")
			cmds, err := pipe.Exec()
			Expect(err).NotTo(HaveOccurred())
			Expect(cmds).To(HaveLen(1))

			val, err := get.Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("value"))
		})

		It("supports custom command", func() {
			pipe.Do("ping")
			cmds, err := pipe.Exec()
			Expect(err).NotTo(HaveOccurred())
			Expect(cmds).To(HaveLen(1))
		})
	}

	Describe("Pipeline", func() {
		BeforeEach(func() {
			pipe = client.Pipeline().(*Pipeline)
		})

		assertPipeline()
	})

	Describe("TxPipeline", func() {
		BeforeEach(func() {
			pipe = client.TxPipeline().(*Pipeline)
		})

		assertPipeline()
	})
})
