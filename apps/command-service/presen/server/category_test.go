package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/adapter"
	"go.uber.org/fx"
)

// Ordered is intentional here (unlike the adapter unit tests):
// there are integration specs where a category created in one spec is
// reused by later specs (Create -> Update -> Delete),
// so they must run in sequence and share state.
var _ = Describe("categoryServer", Ordered, Label("integration-test", "category-server"), func() {
	var (
		server    pbv1.CategoryCommandServer
		container *fx.App
		ctx       context.Context
		category  *pbv1.Category
	)

	// BeforeAll/AfterAll (paired, not BeforeEach/AfterEach) start and stop
	// the fx container exactly once for this ordered group of specs.
	BeforeAll(func() {
		ctx = context.Background()

		container = fx.New(
			application.SrvDepend,
			fx.Provide(
				adapter.NewCategoryAdapterImpl,
				NewCategoryServer,
			),
			fx.Populate(&server),
		)
		Expect(container.Start(ctx)).To(Succeed())
	})

	AfterAll(func() {
		if category != nil {
			_, _ = server.Delete(ctx, &pbv1.CategoryUpParam{
				Crud: pbv1.CRUD_DELETE,
				Id:   category.GetId(),
				Name: category.GetName(),
			})
		}

		Expect(container.Stop(context.Background())).To(Succeed())
	})

	Describe("Create", Label("Create"), func() {
		When("the category does not exist yet", func() {
			It("creates the category and returns no error", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_INSERT,
					Id:   "",
					Name: "Beverages",
				}

				result, err := server.Create(ctx, param)
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Error).To(BeNil())

				category = result.Category
			})
		})

		When("the category is already registered", func() {
			It("returns a CategoryUpResult wrapping a CRUD error", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_INSERT,
					Id:   category.GetId(),
					Name: category.GetName(),
				}

				result, err := server.Create(ctx, param)
				Expect(err).NotTo(HaveOccurred())

				expected := &pbv1.Error{Type: "CRUD Error", Message: "Beverages already exists"}
				Expect(result.Error).To(Equal(expected))
			})
		})
	})

	Describe("Update", Label("Update"), func() {
		When("the category exists", func() {
			It("updates the category and returns no error", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   category.GetId(),
					Name: "Apparel",
				}

				result, err := server.Update(ctx, param)
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Error).To(BeNil())
			})
		})

		When("the category does not exist", func() {
			It("returns a CategoryUpResult wrapping a CRUD error", func() {
				// Generating a fresh random UUID instead of hardcoding one.
				// A hardcoded ID like "...b5c1" only works as long as no
				// other spec, seed data, or DB state happens to create
				// that exact raw. A freshly generated UUID is guaranteed
				// (for all practical purposes) not to collide with any
				// existing raw, so "does not exist" is true by
				// construction rather than by concience.
				id := uuid.NewString()

				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_UPDATE,
					Id:   id,
					Name: "Apparel",
				}

				result, err := server.Update(ctx, param)
				Expect(err).NotTo(HaveOccurred())

				expected := &pbv1.Error{
					Type:    "CRUD Error",
					Message: fmt.Sprintf("Category ID: %s does not exist and could not be updated", id),
				}
				Expect(result.Error).To(Equal(expected))
			})
		})
	})

	Describe("Delete", Label("Delete"), func() {
		When("the category exists", func() {
			It("deletes the category and returns no error", func() {
				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_DELETE,
					Id:   category.GetId(),
					Name: category.GetName(),
				}

				result, err := server.Delete(ctx, param)
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Error).To(BeNil())
			})
		})

		When("the category does not exist", func() {
			It("returns a CategoryUpResult wrapping a CRUD error", func() {
				// Same reasoning as the Update spec above:
				// use a freshly generated UUID rather than a hardcoded "knwon absent" ID,
				// so the spec doesn't depend on today's DB contents.
				id := uuid.NewString()

				param := &pbv1.CategoryUpParam{
					Crud: pbv1.CRUD_DELETE,
					Id:   id,
					Name: "Apparel",
				}

				result, err := server.Delete(ctx, param)
				Expect(err).NotTo(HaveOccurred())

				expected := &pbv1.Error{
					Type:    "CRUD Error",
					Message: fmt.Sprintf("Category ID: %s does not exist and could not be deleted", id),
				}
				Expect(result.Error).To(Equal(expected))
			})
		})
	})
})
