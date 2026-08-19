package adapter

import (
	"github.com/tortillaproduction/go-microservices/apps/query-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/query-service/errs"
	"github.com/tortillaproduction/go-microservices/apps/query-service/infra/gorm/models"
)

type CategoryAdapterImpl struct{}

func NewCategoryAdapterImpl() categories.CategoryAdapter {
	return &CategoryAdapterImpl{}
}

// Category entity -> GOEM model
func (ins *CategoryAdapterImpl) Convert(source *categories.Category) any {
	return &models.Category{
		Id:   source.Id(),
		Name: source.Name(),
	}
}

// GOEM model -> Category entity
func (ins *CategoryAdapterImpl) Rebuild(source any) (dest *categories.Category, err error) {
	if c, ok := source.(*models.Category); ok {
		dest = categories.NewCategory(c.Id, c.Name)
	} else {
		err = errs.NewInternalError("invalid type: expected *models.Category")
	}
	return
}
