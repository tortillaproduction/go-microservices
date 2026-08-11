package products

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("Value Objects constituting Product entity", Ordered, Label("creation of the ProductId struct"), func() {
	var empty_str *errs.DomainError
	var length_over *errs.DomainError
	var not_uuid *errs.DomainError
	var product_id *ProductId
	var uid string

	BeforeAll(func() {
		_, empty_str = NewProductId("")
		_, length_over = NewProductId("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbccccccccccccccccccccdddddddddddddddddd")
		_, not_uuid = NewProductId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		product_id, _ = NewProductId(id.String())
	})

	Context("character count verification", Label("character count"), func() {
		It("if the string is empty, errs.DomainError is returned", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("productID must be 36 characters long")))
		})
		It("if the strings longer than 36 characters, errs.DomainError is returned", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("productID must be 36 characters long")))
		})
	})

	Context("UUID format verification", Label("UUID format"), func() {
		It("if the strings format other than UUIDs, errs.DomainError is returned", func() {
			Expect(not_uuid).To(Equal(errs.NewDomainError("productID must be UUID format")))
		})
		It("if 36-character UUID strings, ProductId is returned", func() {
			id, _ := NewProductId(uid)
			Expect(product_id).To(Equal(id))
		})
	})
})

var _ = Describe("Value Objects constituting Product entity", Ordered, Label("creation of the ProductName struct"), func() {
	var empty_str *errs.DomainError
	var length_over *errs.DomainError
	var product_name *ProductName

	BeforeAll(func() {
		_, empty_str = NewProductName("")
		_, length_over = NewProductName("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbcccccccccccccd")
		product_name, _ = NewProductName("water-based ballpoint pen")
	})

	Context("character count verification", Label("invalid character count"), func() {
		It("if the string is empty, errs.DomainError is returned", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("productName must be between 5 and 30 characters long")))
		})
		It("if the strings longer than 30 characters, errs.DomainError is returned", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("productName must be between 5 and 30 characters long")))
		})
	})

	Context("valid character count verification", Label("valid character count"), func() {
		It("if the strings between 5 and 30 characters, ProductName is returned", func() {
			Expect(product_name.Value()).To(Equal("water-based ballpoint pen"))
		})
	})
})

var _ = Describe("Value Objects constituting Product entity", Ordered, Label("creation of the ProductPrice struct"), func() {
	var min_err *errs.DomainError
	var max_err *errs.DomainError
	var product_price *ProductPrice

	BeforeAll(func() {
		_, min_err = NewProductPrice(49)
		_, max_err = NewProductPrice(10001)
		product_price, _ = NewProductPrice(1500)
	})

	Context("validation of price outside the range", Label("invalid range"), func() {
		It("if the price is less than 50, errs.DomainError is returned", func() {
			Expect(min_err).To(Equal(errs.NewDomainError("productPrice must be between 50 and 10000")))
		})
		It("if the price is more than 10000, errs.DomainError is returned", func() {
			Expect(max_err).To(Equal(errs.NewDomainError("productPrice must be between 50 and 10000")))
		})
	})

	Context("validation of price within the range", Label("valid range"), func() {
		It("if the price is between 50 and 10000, ProductPrice is returned", func() {
			Expect(product_price.Value()).To(Equal(uint32(1500)))
		})
	})
})
