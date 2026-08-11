package categories

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("Category entity", Ordered, Label("creation of the Category struct"), func() {
	Context("instance creation", Label("create Category"), func() {
		It("instantiating a new Category", Label("NewCategory"), func() {
			category_name, _ := NewCategoryName("grocery")
			category, _ := NewCategory(category_name)
			Expect(category.Id()).ToNot(BeNil())
			Expect(category.Name().Value()).To(Equal("grocery"))
		})
		It("rebuild an instance of Category", Label("BuildCategory"), func() {
			category_id, _ := NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			category_name, _ := NewCategoryName("grocery")
			category := BuildCategory(category_id, category_name)
			Expect(category.Id().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			Expect(category.Name().Value()).To(Equal("grocery"))
		})
	})
})

var _ = Describe("Category entity", Ordered, Label("verification of Category equivalency"), func() {
	It("compariosn target is nil", Label("nil verification"), func() {
		By("instance creation")
		category := BuildCategory(nil, nil)
		By("specify nil for Equals()")
		_, err := category.Equals(nil)
		By("evaluate errs.DomainError")
		Expect(err).To(Equal(errs.NewDomainError("nil specified as an argument")))
	})
	It("different CategoryId", Label("false verification"), func() {
		By("creating 2 instances")
		category_name, _ := NewCategoryName("grocery")
		category_a, _ := NewCategory(category_name)
		category_name, _ = NewCategoryName("grocery")
		category_b, _ := NewCategory(category_name)
		By("specify category_b for Equals()")
		result, _ := category_a.Equals(category_b)
		By("evaluate false")
		Expect(result).To(Equal(false))
	})
	It("same CategoryId", Label("true verification"), func() {
		By("creating 2 instances")
		category_name, _ := NewCategoryName("grocery")
		category_a, _ := NewCategory(category_name)

		category_b := BuildCategory(category_a.Id(), category_a.Name())
		By("specify category_b for Equals()")
		result, _ := category_a.Equals(category_b)
		By("evaluate true")
		Expect(result).To(Equal(true))
	})
})
