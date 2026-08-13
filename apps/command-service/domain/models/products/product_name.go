package products

import (
	"fmt"
	"unicode/utf8"

	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// value object
type ProductName struct {
	value string
}

func (ins *ProductName) Value() string {
	return ins.value
}

func NewProductName(value string) (*ProductName, *errs.DomainError) {
	const MIN_LENGTH int = 5
	const MAX_LENGTH int = 100

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("productName must be between %d and %d characters long", MIN_LENGTH, MAX_LENGTH))
	}

	return &ProductName{value: value}, nil
}
