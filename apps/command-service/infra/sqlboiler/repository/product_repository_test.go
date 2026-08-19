package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("productRepositorySQLBoiler", Ordered, Label("unit-test", "product-repository"), func() {
	var category *categories.Category
	var rep products.ProductRepository
	var ctx context.Context
	var tran *sql.Tx

	BeforeAll(func() {
		rep = NewProductRepositorySQLBoiler()
		category_ID, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		category_Name, _ := categories.NewCategoryName("Stationery")
		category = categories.BuildCategory(category_ID, category_Name)
	})

	BeforeEach(func() {
		ctx = context.Background()
		var err error

		tran, err = boil.BeginTx(ctx, nil)
		Expect(err).NotTo(HaveOccurred(), "Failed to bigin transaction - check DB connection")
		Expect(tran).NotTo(BeNil())
	})

	AfterEach(func() {
		tran.Rollback()
	})

	// Tests for Exists() method
	Describe("Exists", Label("Exists"), func() {
		Context("when the product does not exist", func() {
			It("returns nil", func() {
				productID, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e500")
				productName, _ := products.NewProductName("Water-based Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(100))
				product := products.BuildProduct(productID, productName, productPrice, category)
				result := rep.Exists(ctx, tran, product)
				Expect(result).To(BeNil())
			})
		})

		Context("when a product with the same name already exists", func() {
			It("returns a CRUDError", func() {
				productName, _ := products.NewProductName("Water-based Ballpoint Pen (Black)")
				product := products.BuildProduct(nil, productName, nil, nil)

				result := rep.Exists(ctx, tran, product)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("%s already exists", product.Name().Value()))))
			})
		})
	})

	// Tests for Create() method
	Describe("Create", Label("Create"), func() {
		Context("when creating a new product", func() {
			It("successfully creates the product and returns nil", func() {
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(150))
				product, _ := products.NewProduct(productName, productPrice, category)
				result := rep.Create(ctx, tran, product)
				Expect(result).To(BeNil())
			})
		})

		Context("when adding a product with a duplicate id", func() {
			It("returns a CRUDError for unique constraint violation", func() {
				productID, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(200))
				product := products.BuildProduct(productID, productName, productPrice, category)
				result := rep.Create(ctx, tran, product)
				Expect(result).To(Equal(errs.NewCRUDError("unique constraint violation")))
			})
		})
	})

	// Tests for UpdateById() method
	Describe("UpdateById", Label("UpdateById"), func() {
		Context("when the specified id does not exist", func() {
			It("returns a CRUDError", func() {
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(200))
				product, _ := products.NewProduct(productName, productPrice, category)
				result := rep.UpdateById(ctx, tran, product)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("Product ID: %s does not exist and could not be updated", product.Id().Value()),
				)))
			})
		})

		Context("when the specified id exists", func() {
			It("successfully updates the product and returns nil", func() {
				productID, _ := products.NewProductId("8f81a72a-58ef-422b-b472-d982e8665292")
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(200))
				product := products.BuildProduct(productID, productName, productPrice, category)
				result := rep.UpdateById(ctx, tran, product)
				Expect(result).To(BeNil())
			})
		})
	})

	// Tests for DeleteById() method
	Describe("DeleteById", Label("DeleteById"), func() {
		Context("when the specified id does not exist", func() {
			It("returns a CRUDError", func() {
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(200))
				product, _ := products.NewProduct(productName, productPrice, category)
				result := rep.DeleteById(ctx, tran, product)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("Product ID: %s does not exist and could not be deleted", product.Id().Value()),
				)))
			})
		})

		Context("when the specified id exists", func() {
			It("successfully deletes the product and returns nil", func() {
				productID, _ := products.NewProductId("8f81a72a-58ef-422b-b472-d982e8665292")
				productName, _ := products.NewProductName("Ballpoint Pen")
				productPrice, _ := products.NewProductPrice(uint32(200))
				product := products.BuildProduct(productID, productName, productPrice, category)
				rep.Create(ctx, tran, product)

				result := rep.DeleteById(ctx, tran, product)
				Expect(result).To(BeNil())
			})
		})
	})
})
