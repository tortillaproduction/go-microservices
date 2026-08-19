package products

import "github.com/tortillaproduction/go-microservices/apps/query-service/domain/models/categories"

// Entity
type Product struct {
	id       string
	name     string
	price    uint32
	category *categories.Category
}

func NewProduct(id string, name string, price uint32, category *categories.Category) *Product {
	return &Product{id: id, name: name, price: price, category: category}
}

func (ins *Product) Id() string {
	return ins.id
}

func (ins *Product) Name() string {
	return ins.name
}

func (ins *Product) Price() uint32 {
	return ins.price
}

func (ins *Product) Category() *categories.Category {
	return ins.category
}
