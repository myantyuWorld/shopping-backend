package repository_test

import (
	"context"

	domainModel "github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	domainRepo "github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository"
	dbModel "github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShoppingItemRepository", func() {
	var repo domainRepo.IShoppingItemRepository
	Context("test data input", func() {
		items := []dbModel.ShoppingItem{
			{
				OwnerID:  0,
				Category: "food",
				Name:     "food a",
				Picked:   false,
			},
			{
				OwnerID:  1,
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
			{
				OwnerID:  1,
				Category: "necessity",
				Name:     "necessity a",
				Picked:   false,
			},
		}
		BeforeEach(func() {
			repo = repository.NewShoppingItemRepository(database)
		})
		It("fixture", func() {
			Expect("hello").NotTo(BeNil())
			res := database.Create(items)
			Expect(res.Error).To(BeNil())
		})
	})
	Context("testing", func() {
		It("Find", func() {
			res, err := repo.FindByOwnerID(context.TODO(), 0)
			Expect(res).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(len(res)).To(Equal(2))
		})
		It("Save", func() {
			err := repo.Save(context.TODO(), &domainModel.ShoppingItem{
				OwnerID:  0,
				Category: "food",
				Name:     "save 1",
				Picked:   false,
			})
			Expect(err).To(BeNil())
		})
		It("Logical Delete", func() {
			res, err := repo.FindByOwnerID(context.TODO(), 0)
			Expect(res).NotTo(BeNil())
			Expect(err).To(BeNil())

			err2 := repo.LogicalDelete(context.TODO(), res[0].ID)
			Expect(err2).To(BeNil())
		})
	})
})
