package categories

import (
	"fmt"
	"unicode/utf8"

	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// value object
type CategoryName struct {
	value string
}

func (ins *CategoryName) Value() string {
	return ins.value
}

func NewCategoryName(value string) (*CategoryName, *errs.DomainError) {
	const MIN_LENGTH int = 2
	const MAX_LENGTH int = 20

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("categoryName must be between %d and %d characters long", MIN_LENGTH, MAX_LENGTH))
	}

	return &CategoryName{value: value}, nil
}
