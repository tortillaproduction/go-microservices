package products

import (
	"github.com/google/uuid"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// entity
type Product struct {
	id       *ProductId
	name     *ProductName
	price    *ProductPrice
	category *categories.Category
}

func (ins *Product) Id() *ProductId {
	return ins.id
}

func (ins *Product) Name() *ProductName {
	return ins.name
}

func (ins *Product) Price() *ProductPrice {
	return ins.price
}

func (ins *Product) Category() *categories.Category {
	return ins.category
}

// change value
func (ins *Product) ChangeProductName(name *ProductName) {
	ins.name = name
}

func (ins *Product) ChangeProductPrice(price *ProductPrice) {
	ins.price = price
}

func (ins *Product) ChangeCategory(category *categories.Category) {
	ins.category = category
}

// equivalence check
// entity verify equivalence using an ID
func (ins *Product) Equals(obj *Product) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("nil specified as an argument")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

func NewProduct(
	name *ProductName,
	price *ProductPrice,
	category *categories.Category,
) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewProductId(uid.String()); err != nil {
			return nil, err
		} else {
			return &Product{
				id:       id,
				name:     name,
				price:    price,
				category: category,
			}, nil
		}
	}
}

// re-build entity
// used when retrieving a value from database and regenerating an entity using that value
func BuildProduct(
	id *ProductId,
	name *ProductName,
	price *ProductPrice,
	category *categories.Category,
) *Product {
	product := Product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
	return &product
}
