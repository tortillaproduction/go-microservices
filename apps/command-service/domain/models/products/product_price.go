package products

import (
	"fmt"

	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// value object
type ProductPrice struct {
	value uint32
}

func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

func NewProductPrice(value uint32) (*ProductPrice, *errs.DomainError) {
	const MIN_VALUE = 50
	const MAX_VALUE = 10000

	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(fmt.Sprintf("productPrice must be between %d and %d", MIN_VALUE, MAX_VALUE))
	}

	return &ProductPrice{value: value}, nil
}
