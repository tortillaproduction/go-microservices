package adapter

import (
	"github.com/tortillaproduction/go-microservices/apps/query-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/query-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/query-service/errs"
	"github.com/tortillaproduction/go-microservices/apps/query-service/infra/gorm/models"
)

type ProductAdapterImpl struct{}

func NewProductAdapterImpl() products.ProductAdapter {
	return &ProductAdapterImpl{}
}

// Product entity -> GORM model
func (ins *ProductAdapterImpl) Convert(source *products.Product) any {
	return &models.Product{
		Id:           source.Id(),
		Name:         source.Name(),
		Price:        source.Price(),
		CategoryId:   source.Category().Id(),
		CategoryName: source.Category().Name(),
	}
}

// GORM model -> Product entity
func (ins *ProductAdapterImpl) Rebuild(source any) (dest *products.Product, err error) {
	if p, ok := source.(*models.Product); ok {
		c := categories.NewCategory(p.CategoryId, p.CategoryName)
		dest = products.NewProduct(p.Id, p.Name, p.Price, c)
	} else {
		err = errs.NewInternalError("invalid type: expected *models.Product")
	}
	return
}
