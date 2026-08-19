package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("categoryRepositorySQLBoiler", Ordered, Label("unit-test", "category-repository"), func() {
	var rep categories.CategoryRepository
	var ctx context.Context
	var tran *sql.Tx

	BeforeAll(func() {
		rep = NewCategoryRepositorySQLBoiler()
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
		Context("when the category does not exist", func() {
			It("returns nil", func() {
				name, _ := categories.NewCategoryName("Grocery")
				category, _ := categories.NewCategory(name)
				result := rep.Exists(ctx, tran, category)
				Expect(result).To(BeNil())
			})
		})

		Context("when a category with the sama name already exists", func() {
			It("returns a CRUDError", func() {
				name, _ := categories.NewCategoryName("Stationery")
				category, _ := categories.NewCategory(name)
				result := rep.Exists(ctx, tran, category)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("%s already exists", category.Name().Value()),
				)))
			})
		})
	})

	// Tests for Create() method
	Describe("Create", Label("Create"), func() {
		Context("when creating a new product category", func() {
			It("successfully creates the category and returns nil", func() {
				name, _ := categories.NewCategoryName("Grocery")
				category, _ := categories.NewCategory(name)
				result := rep.Create(ctx, tran, category)
				Expect(result).To(BeNil())
			})
		})

		Context("when adding a category with a duplicate id", func() {
			It("returns a CRUDError for unique constraint violation", func() {
				id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				name, _ := categories.NewCategoryName("Stationery")
				category := categories.BuildCategory(id, name)
				result := rep.Create(ctx, tran, category)
				Expect(result).To(Equal(errs.NewCRUDError("unique constraint violation")))
			})
		})
	})

	// Tests for UpdateById() method
	Describe("UpdateById", Label("UpdateById"), func() {
		Context("when the specified id does not exist", func() {
			It("returns a CRUDError", func() {
				id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
				name, _ := categories.NewCategoryName("Stationery")
				category := categories.BuildCategory(id, name)
				result := rep.UpdateById(ctx, tran, category)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("Category ID: %s does not exist and could not be updated", category.Id().Value()),
				)))
			})

			Context("when the specified id exists", func() {
				It("successfully updates the category and returns nil", func() {
					id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
					name, _ := categories.NewCategoryName("Stationery 1")
					category := categories.BuildCategory(id, name)
					result := rep.UpdateById(ctx, tran, category)
					Expect(result).To(BeNil())
				})
			})
		})
	})

	// Tests for DeleteById() method
	Describe("DeleteById", Label("DeleteById"), func() {
		Context("when the specified id does not exist", func() {
			It("returns a CRUDError", func() {
				id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
				name, _ := categories.NewCategoryName("Stationery 1")
				category := categories.BuildCategory(id, name)
				result := rep.DeleteById(ctx, tran, category)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("Category ID: %s does not exist and could not be deleted", category.Id().Value()),
				)))
			})

			Context("when the specified id exists", func() {
				It("successfully deletes the category and returns nil", func() {
					name, _ := categories.NewCategoryName("Grocery")
					category, _ := categories.NewCategory(name)
					rep.Create(ctx, tran, category)

					result := rep.DeleteById(ctx, tran, category)
					Expect(result).To(BeNil())
				})
			})
		})
	})
})
