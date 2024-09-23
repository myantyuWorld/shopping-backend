package interacter_test

import (
	"context"
	"errors"

	domainModel "github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	repoMock "github.com/LeoTwins/go-clean-architecture/internal/domain/repository/mock"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/interacter"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShoppingItemInteractor", func() {
	var ctrl *gomock.Controller

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("TBD", func() {
		BeforeEach(func() {})
		JustBeforeEach(func() {})
		Context("", func() {
			BeforeEach(func() {})
			It("", func() {})
		})
		Context("", func() {
			BeforeEach(func() {})
			It("", func() {})
		})
	})
	Describe("Remove", func() {
		var (
			interactor input.IShoppingItemUsecase
			err        error
			itemID     uint
		)
		BeforeEach(func() {
			presenter := mock.NewMockIShoppingItemPresenter(ctrl)
			mockRepo := repoMock.NewMockIShoppingItemRepository(ctrl)
			mockRepo.EXPECT().LogicalDelete(gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, itemID uint) error {
					switch itemID {
					case 1:
						return nil
					default:
						return errors.New("")
					}
				}).AnyTimes()
			interactor = interacter.NewShoppingItemUsecase(mockRepo, presenter)
		})
		JustBeforeEach(func() {
			err = interactor.Remove(context.TODO(), itemID)
		})
		Context("", func() {
			BeforeEach(func() {
				itemID = 1
			})
			It("", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			BeforeEach(func() {
				itemID = 999
			})
			It("", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Register", func() {
		var (
			interactor input.IShoppingItemUsecase
			err        error
			ownerID    uint
			category   string
			name       string
		)
		BeforeEach(func() {
			presenter := mock.NewMockIShoppingItemPresenter(ctrl)
			mockRepo := repoMock.NewMockIShoppingItemRepository(ctrl)
			mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, item *domainModel.ShoppingItem) error {
					switch ownerID {
					case 1:
						return nil
					default:
						return errors.New("")
					}
				}).AnyTimes()
			interactor = interacter.NewShoppingItemUsecase(mockRepo, presenter)
		})
		JustBeforeEach(func() {
			err = interactor.Register(context.TODO(), ownerID, name, category)
		})
		Context("valid", func() {
			BeforeEach(func() {
				ownerID = 1
				name = "aaaa"
				category = "food"
			})
			It("", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("invalid(repository)", func() {
			BeforeEach(func() {
				ownerID = 999
				name = "aaaa"
				category = "food"
			})
			It("", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Find", func() {
		var (
			interactor    input.IShoppingItemUsecase
			ownerID       uint
			output        []*dto.ShoppingItemOutput
			err           error
			noDataOwnerID uint
		)
		noDataOwnerID = 999
		BeforeEach(func() {
			presenter := mock.NewMockIShoppingItemPresenter(ctrl)
			presenter.EXPECT().Output(gomock.Any()).DoAndReturn(
				func(item interface{}) interface{} {
					return []*dto.ShoppingItemOutput{
						{
							ID:       0,
							OwnerID:  ownerID,
							Name:     "food",
							Category: "aaa",
						},
						{
							ID:       1,
							OwnerID:  ownerID,
							Name:     "food",
							Category: "bbb",
						},
					}
				}).AnyTimes()
			mockRepo := repoMock.NewMockIShoppingItemRepository(ctrl)
			mockRepo.EXPECT().FindByOwnerID(gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, ownerID uint) ([]*domainModel.ShoppingItem, error) {
					switch ownerID {
					case noDataOwnerID:
						return []*domainModel.ShoppingItem{}, nil
					default:
						return []*domainModel.ShoppingItem{
							{
								ID:       0,
								OwnerID:  ownerID,
								Category: "food",
								Name:     "aaa",
								Picked:   false,
							},
							{
								ID:       1,
								OwnerID:  ownerID,
								Category: "food",
								Name:     "bbb",
								Picked:   false,
							},
						}, nil
					}
				}).AnyTimes()
			interactor = interacter.NewShoppingItemUsecase(mockRepo, presenter)
		})
		JustBeforeEach(func() {
			output, err = interactor.Find(context.TODO(), ownerID)
		})
		Context("valid response", func() {
			BeforeEach(func() {
				ownerID = 1
			})
			It("", func() {
				Expect(output).NotTo(BeNil())
				Expect(len(output)).To(Equal(2))
				Expect(err).To(BeNil())
			})
		})
		Context("no data", func() {
			BeforeEach(func() {
				ownerID = noDataOwnerID
			})
			It("", func() {
				Expect(output).NotTo(BeNil())
				Expect(len(output)).To(Equal(0))
				Expect(err).To(BeNil())
			})
		})
	})
})
