package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShoppingItemRepository", func() {
	Context("test data input", func() {
		BeforeEach(func() {
		})
		It("fixture", func() {
			Expect("hello").NotTo(BeNil())
		})
	})
})
