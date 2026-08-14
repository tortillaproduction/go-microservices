package adapter

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

var _ = Describe("categoryAdapterImpl", Label("unit-test", "category-adapter"), func() {
	var adapter CategoryAdapter

	// BeforeEach (instead of BeforeAll) keeps every spec independent and
	// prevents state from one test leaking into another.
	BeforeEach(func() {
		adapter = NewCategoryAdapterImpl()
	})

	// Tests for ToEntity() method
	Describe("ToEntity", Label("ToEntity"), func() {
		When("both Id and Name are provided", func() {
			It("returns an entity.Category with both fields set", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name: "Stationary",
				}

				result, err := adapter.ToEntity(param)
				Expect(err).NotTo(HaveOccurred())

				id, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				name, err := categories.NewCategoryName("Stationary")
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(Equal(categories.BuildCategory(id, name)))
			})
		})

		When("only Id is provided", func() {
			It("returns an entity.Category with a nil Name", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_DELETE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name: "",
				}

				result, err := adapter.ToEntity(param)
				Expect(err).NotTo(HaveOccurred())

				id, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(Equal(categories.BuildCategory(id, nil)))
			})
		})

		When("only Name is provided", func() {
			It("returns an entity.Category with the Name field set", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_INSERT,
					Id:   "",
					Name: "Stationary",
				}

				result, err := adapter.ToEntity(param)
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Name().Value()).To(Equal("Stationary"))
			})
		})

		// A DescribeTable collapses the five near-identical "invalid input,
		// expect a matching DomainError" specs into a single, easy-to-extend
		// table.
		DescribeTable(
			"when given invalid input",
			func(param *pbv1.CategoryUpParam, expectedMessage string) {
				_, err := adapter.ToEntity(param)
				Expect(err).To(Equal(errs.NewDomainError(expectedMessage)))
			},

			Entry(
				"an Id shorter than 36 characters",
				&pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c",
					Name: "Stationary",
				},
				"categoryID must be 36 characters long",
			),
			Entry(
				"an Id longer than 36 characters",
				&pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0ac",
					Name: "Stationary",
				},
				"categoryID must be 36 characters long",
			),
			Entry(
				"an Id that is not in UUID format",
				&pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "aaaaaaaaaabbbbbbbbbbccccccccccdddddd",
					Name: "Stationary",
				},
				"categoryID must be UUID format",
			),
			Entry(
				"a Name shorter than 2 characters",
				&pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name: "S",
				},
				"categoryName must be between 2 and 20 characters long",
			),
			Entry(
				"a Name longer than 20 characters",
				&pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name: "Stationary Stationary Stationary",
				},
				"categoryName must be between 2 and 20 characters long",
			),
		)
	})

	// Tests for ToResult() method
	Describe("ToResult", Label("ToResult"), func() {
		When("given an entity.Category", func() {
			It("returns a CategoryUpResult wrapping the corresponding pbv1.Category", func() {
				id, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				name, err := categories.NewCategoryName("Stationary")
				Expect(err).NotTo(HaveOccurred())

				category := categories.BuildCategory(id, name)
				result := adapter.ToResult(category)

				expected := &pbv1.Category{Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: "Stationary"}
				Expect(result.Category).To(Equal(expected))
			})
		})

		// A DescribeTable collapses three near-identical "given an error, expect
		// a matching pbv1.Error" specs into a single, easy-to-extend table.
		// Note: expected is passed as *pbv1.Error (not pbv1.Error) to avoid copying
		// the generated protobuf struct by value, which go vet flags because it
		// embeds protoimpl.MessageState (containing a sync.Mutex).
		DescribeTable(
			"when given an error",
			func(inputErr error, expected *pbv1.Error) {
				result := adapter.ToResult(inputErr)
				Expect(result.Error).To(Equal(expected))
			},

			Entry(
				"DomainError",
				errs.NewDomainError("Stationary is already registered."),
				&pbv1.Error{Type: "Domain Error", Message: "Stationary is already registered."},
			),

			Entry(
				"CRUDError",
				errs.NewCRUDError("Stationary is already registered."),
				&pbv1.Error{Type: "CRUD Error", Message: "Stationary is already registered."},
			),
			Entry(
				"InternalError",
				errs.NewInternalError("Stationary is already registered."),
				&pbv1.Error{Type: "Internal Error", Message: "Internal Server Error"},
			),
		)
	})
})
