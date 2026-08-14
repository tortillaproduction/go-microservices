package adapter

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("productAdapterImpl", Label("unit-test", "product-adapter"), func() {
	var (
		adapter  ProductAdapter
		category *categories.Category
	)

	// BeforeEach (instead of BeforeAll) keeps every spec independent and
	// prevents state from one test leaking into another.
	BeforeEach(func() {
		adapter = NewProductAdapterImpl()

		categoryId, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		Expect(err).NotTo(HaveOccurred())

		category = categories.BuildCategory(categoryId, nil)
	})

	// Tests for ToEntity() method
	Describe("ToEntity", Label("ToEntity"), func() {
		When("the CRUD operation is UPDATE and an Id is provided", func() {
			It("maps Id, Name, and CategoryId onto the entity", func() {
				param := &pbv1.ProductUpParam{
					Crud:       pbv1.CRUD_UPDATE,
					Id:         "ac413f22-0cf1-490a-9635-7e9ca810e544",
					Name:       "Water-based Ballpoint Pen (Black)",
					Price:      120,
					CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				}

				result, err := adapter.ToEntity(param)
				Expect(err).NotTo(HaveOccurred())

				prductId, err := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				Expect(err).NotTo(HaveOccurred())
				productName, err := products.NewProductName("Water-based Ballpoint Pen (Black)")
				Expect(err).NotTo(HaveOccurred())
				productPrice, err := products.NewProductPrice(uint32(120))
				Expect(err).NotTo(HaveOccurred())

				expected := products.BuildProduct(prductId, productName, productPrice, category)
				Expect(result).To(Equal(expected))
			})
		})

		When("the CRUD operation is INSERT and no Id is provided", func() {
			It("generates a new Id and maps Name, Price and CategoryId onto the entity", func() {
				param := &pbv1.ProductUpParam{
					Crud:       pbv1.CRUD_INSERT,
					Id:         "",
					Name:       "Water-based Ballpoint Pen (Black)",
					Price:      120,
					CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				}

				result, err := adapter.ToEntity(param)
				Expect(err).NotTo(HaveOccurred())

				Expect(result.Id()).NotTo(BeNil())
				Expect(result.Name().Value()).To(Equal("Water-based Ballpoint Pen (Black)"))
				Expect(result.Price().Value()).To(Equal(uint32(120)))
				Expect(result.Category().Id().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			})
		})
	})

	// Tests for ToResult() method
	Describe("ToResult", Label("ToResult"), func() {
		When("given a products.Product entity", func() {
			It("returns a ProductResult wrapping the corresponding pbv1.Product", func() {
				prductId, err := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				Expect(err).NotTo(HaveOccurred())
				productName, err := products.NewProductName("Water-based Ballpoint Pen (Black)")
				Expect(err).NotTo(HaveOccurred())
				productPrice, err := products.NewProductPrice(uint32(120))
				Expect(err).NotTo(HaveOccurred())

				product := products.BuildProduct(prductId, productName, productPrice, category)

				result := adapter.ToResult(product)

				expectedCategory := pbv1.Category{Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: ""}
				expectedProduct := pbv1.Product{
					Id:       "ac413f22-0cf1-490a-9635-7e9ca810e544",
					Name:     "Water-based Ballpoint Pen (Black)",
					Price:    120,
					Category: &expectedCategory,
				}

				Expect(result.Product).To(Equal(&expectedProduct))
				Expect(result.Error).To(BeNil())
			})
		})

		// A DescribeTable collapses three near-idential "guven an error, expect
		// a matching pbv1.Error" specs into a single, easy-to-extend table.
		DescribeTable(
			"when given an error", func(inputErr error, expected *pbv1.Error) {
				result := adapter.ToResult(inputErr)
				Expect(result.Error).To(Equal(expected))
			},

			Entry(
				"DomainError",
				errs.NewDomainError("Water-based Ballpoint Pen (Black) is already registered."),
				&pbv1.Error{Type: "Domain Error", Message: "Water-based Ballpoint Pen (Black) is already registered."},
			),

			Entry(
				"CRUDError",
				errs.NewCRUDError("Water-based Ballpoint Pen (Black) is already registered."),
				&pbv1.Error{Type: "CRUD Error", Message: "Water-based Ballpoint Pen (Black) is already registered."},
			),

			Entry(
				"InternalError",
				errs.NewInternalError("Internal Server Error"),
				&pbv1.Error{Type: "Internal Error", Message: "Internal Server Error"},
			),
		)
	})
})
