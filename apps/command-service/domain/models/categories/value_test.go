package categories

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("Value objects constituting Category entity", Ordered, Label("creation of the CategoryId struct"), func() {
	var empty_str *errs.DomainError
	var length_over *errs.DomainError
	var not_uuid *errs.DomainError
	var category_id *CategoryId
	var uid string

	BeforeAll(func() {
		_, empty_str = NewCategoryId("")
		_, length_over = NewCategoryId("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbccccccccccccccccccccdddddddddddddddddd")
		_, not_uuid = NewCategoryId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		category_id, _ = NewCategoryId(id.String())
	})

	Context("character count verification", Label("character count"), func() {
		It("if the string is empty, errs.DomainError is returned", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("categoryID must be 36 characters long")))
		})
		It("if the strings longer than 36 characters, errs.DomainError is returned", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("categoryID must be 36 characters long")))
		})
	})

	Context("UUID format verification", Label("UUID format"), func() {
		It("if the strings format other than UUIDs, errs.DomainError is returned", func() {
			Expect(not_uuid).To(Equal(errs.NewDomainError("categoryID must be UUID format")))
		})
		It("if 36-character UUID strings, CategoryId is returned", func() {
			Expect(category_id.Value()).To(Equal(uid))
		})
	})

	Context("equality verification", Label("equality"), func() {
		It("returns true if the addresses are equal", func() {
			result := category_id.Equals(category_id)
			Expect(result).To(Equal(true))
		})
		It("returns true if the values are equal", func() {
			c_id, _ := NewCategoryId(uid)
			result := category_id.Equals(c_id)
			Expect(result).To(Equal(true))
		})
		It("returns false if the values are not equal", func() {
			uid, _ := uuid.NewRandom()
			c_id, _ := NewCategoryId(uid.String())
			result := category_id.Equals(c_id)
			Expect(result).To(Equal(false))
		})
	})
})

var _ = Describe("Value objects constituting Category entity", Ordered, Label("creation of the CategoryName struct"), func() {
	var empty_str *errs.DomainError
	var length_over *errs.DomainError
	var category_name *CategoryName

	BeforeAll(func() {
		_, empty_str = NewCategoryName("")
		_, length_over = NewCategoryName("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbcccccccccccccd")
		category_name, _ = NewCategoryName("grocery")
	})

	Context("character count verification", Label("character count"), func() {
		It("if the string is empty, errs.DomainError is returned", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("categoryName must be between 2 and 20 characters long")))
		})
		It("if the strings longer than 20 characters, errs.DomainError is returned", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("categoryName must be between 2 and 20 characters long")))
		})
	})

	Context("valid character count verification", Label("valid character count"), func() {
		It("if the strings between 2 and 20 characters, CategoryName is returned", func() {
			Expect(category_name.Value()).To(Equal("grocery"))
		})
	})
})
