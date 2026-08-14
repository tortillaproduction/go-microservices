package impl_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
	"go.uber.org/fx"
)

var _ = Describe("productServiceImpl", Ordered, Label("unit-test", "product-service"), func() {
	var product *products.Product
	var category *categories.Category
	var service service.ProductService
	var ctx context.Context
	var container *fx.App

	BeforeAll(func() {
		categoryID, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		categoryName, _ := categories.NewCategoryName("Stationery")
		category = categories.BuildCategory(categoryID, categoryName)

		productName, _ := products.NewProductName("Ballpoint Pen")
		productPrice, _ := products.NewProductPrice(uint32(150))
		product, _ = products.NewProduct(productName, productPrice, category)

		ctx = context.Background()

		container = fx.New(
			application.SrvDepend,
			fx.Populate(&service),
		)

		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func() {
		err := container.Stop(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	// Tests for Add() method
	Describe("Add", Label("Add"), func() {
		Context("when adding a new product", func() {
			It("successfully creates the product and returns nil", func() {
				result := service.Add(ctx, product)
				Expect(result).To(BeNil())
			})
		})

		Context("when the product name already exists", func() {
			It("returns a CRUDError", func() {
				result := service.Add(ctx, product)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("%s already exists", product.Name().Value()))))
			})
		})
	})

	// Tests for Update() method
	Describe("Update", Label("Update"), func() {
		Context("when the specified obj_id exists", func() {
			It("successfully updates the product and returns nil", func() {
				productName, _ := products.NewProductName("Ballpoint Pen (Black)")
				productPrice, _ := products.NewProductPrice(uint32(200))
				upProduct := products.BuildProduct(product.Id(), productName, productPrice, category)

				result := service.Update(ctx, upProduct)
				Expect(result).To(BeNil())
			})
		})

		Context("when the specified obj_id does not exist", func() {
			It("returns a CRUDError", func() {
				name, _ := products.NewProductName("Ballpoint Pen (Black)")
				productPrice, _ := products.NewProductPrice(uint32(200))
				upProduct, _ := products.NewProduct(name, productPrice, category)

				result := service.Update(ctx, upProduct)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("Product ID: %s does not exist and could not be updated", upProduct.Id().Value()))))
			})
		})
	})

	// Tests for Delete() method
	Describe("Delete", Label("Delete"), func() {
		Context("when the specified obj_id exists", func() {
			It("successfully deletes the product and returns nil", func() {
				result := service.Delete(ctx, product)
				Expect(result).To(BeNil())
			})
		})

		Context("when the specified obj_id does not exist", func() {
			It("returns a CRUDError", func() {
				productName, _ := products.NewProductName("Ballpoint Pen (Black)")
				productPrice, _ := products.NewProductPrice(uint32(200))
				delProduct, _ := products.NewProduct(productName, productPrice, category)

				result := service.Delete(ctx, delProduct)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("Product ID: %s does not exist and could not be deleted", delProduct.Id().Value()))))
			})
		})
	})
})
