package repository_test

import (
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShoppingItemRepository", func() {
	Context("test data input", func() {
		items := []model.ShoppingItem{
			{
				OwnerID:  0,
				Category: "food",
				Name:     "food a",
				Picked:   false,
			},
			{
				OwnerID:  0,
				Category: "necessity",
				Name:     "necessity a",
				Picked:   false,
			},
		}
		BeforeEach(func() {
		})
		It("fixture", func() {
			Expect("hello").NotTo(BeNil())
			res := database.Create(items)
			Expect(res.Error).To(BeNil())
		})
	})
})
