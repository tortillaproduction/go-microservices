package impl_test

import (
	"context"
	"fmt"
	"log"

	"github.com/tortillaproduction/go-microservices/apps/command-service/application"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("categoryServiceImpl", Ordered, Label("unit-test", "category-service"), func() {
	var category *categories.Category
	var service service.CategoryService
	var ctx context.Context
	var container *fx.App

	BeforeAll(func() {
		name, _ := categories.NewCategoryName("Bevarages")
		category, _ = categories.NewCategory(name)

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
		Context("when adding a new category", func() {
			It("successfully creates the category and returns nil", func() {
				result := service.Add(ctx, category)
				Expect(result).To(BeNil())
			})
		})

		Context("when the category name already exists", func() {
			It("returns a CRUDError", func() {
				result := service.Add(ctx, category)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("%s already exists", category.Name().Value()))))
			})
		})
	})

	// Tests for Update() method
	Describe("Update", Label("Update"), func() {
		Context("when the specified obj_id exists", func() {
			It("successfully updates the category and returns nil", func() {
				result := service.Update(ctx, category)
				log.Println("returns nil when specified obj_id exists:", result)
				Expect(result).To(BeNil())
			})
		})

		Context("when the specified obj_id does not exist", func() {
			It("returns a CRUDError", func() {
				name, _ := categories.NewCategoryName("Bevarages")
				upCategory, _ := categories.NewCategory(name)
				result := service.Update(ctx, upCategory)
				log.Println("returns CRUDError when specified obj_id does not exist:", result)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("Category ID: %s does not exist and could not be updated", upCategory.Id().Value()))))
			})
		})
	})

	// Tests for Delete() method
	Describe("Delete", Label("Delete"), func() {
		Context("when the specified obj_id exists", func() {
			It("successfully deletes the category and returns nil", func() {
				result := service.Delete(ctx, category)
				Expect(result).To(BeNil())
			})
		})

		Context("when the specified obj_id does not exist", func() {
			It("returns a CRUDError", func() {
				name, _ := categories.NewCategoryName("Bevarages")
				delCategory, _ := categories.NewCategory(name)
				result := service.Delete(ctx, delCategory)
				Expect(result).To(Equal(errs.NewCRUDError(fmt.Sprintf("Category ID: %s does not exist and could not be deleted", delCategory.Id().Value()))))
			})
		})
	})
})
