package products

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("Product entity", Ordered, Label("creation of the Product struct"), func() {
	BeforeAll(func() {})

	_ = Describe("Product entity", Ordered, Label("creation of the Product struct"), func() {
		Context("instance creation", Label("Create Product"), func() {
			It("instantiating a new Product", Label("NewProduct"), func() {
				name, _ := NewProductName("chocolate")
				price, _ := NewProductPrice(150)
				product, _ := NewProduct(name, price, nil)
				Expect(product.Id().Value()).ToNot(BeNil())
				Expect(product.Name().Value()).To(Equal("chocolate"))
				Expect(product.Price().Value()).To(Equal(uint32(150)))
				Expect(product.Category()).To(BeNil())
			})

			It("rebuild an instance of Product", Label("BuildProduct"), func() {
				id, _ := NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				name, _ := NewProductName("chocolate")
				price, _ := NewProductPrice(150)
				product := BuildProduct(id, name, price, nil)
				Expect(product.Id().Value()).To(Equal("ac413f22-0cf1-490a-9635-7e9ca810e544"))
				Expect(product.Name().Value()).To(Equal("chocolate"))
				Expect(product.Price().Value()).To(Equal(uint32(150)))
				Expect(product.Category()).To(BeNil())
			})
		})
	})
})

var _ = Describe("Product entity", Ordered, Label("verification of Product equivalency"), func() {
	var category *categories.Category
	var product *Product

	BeforeAll(func() {
		category_name, _ := categories.NewCategoryName("grocery")
		category, _ = categories.NewCategory(category_name)
		product_name, _ := NewProductName("potato chips")
		product_price, _ := NewProductPrice(uint32(200))
		product, _ = NewProduct(product_name, product_price, category)
	})

	Context("verification of error", func() {
		It("comparison target is nil", Label("nil verification"), func() {
			By("verify that specifying nil returns a DomainError")
			_, err := product.Equals(nil)
			Expect(err).To(Equal(errs.NewDomainError("nil specified as an argument")))
		})
	})

	Context("verification of comparison results", func() {
		It("differnt ProductId", Label("false verification"), func() {
			product_name, _ := NewProductName("potato chips")
			product_price, _ := NewProductPrice(uint32(200))
			p, _ := NewProduct(product_name, product_price, category)
			By("verify that specifying different Product returns false")
			result, _ := product.Equals(p)
			Expect(result).To(Equal(false))
		})

		It("same ProductId", Label("true verification"), func() {
			p := BuildProduct(
				product.Id(),
				product.Name(),
				product.Price(),
				category,
			)
			By("verify that specifying same Product returns true")
			result, _ := product.Equals(p)
			Expect(result).To(Equal(true))
		})
	})
})
