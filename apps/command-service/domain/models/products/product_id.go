package products

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// value object (UUID)
type ProductId struct {
	value string
}

func (ins *ProductId) Value() string {
	return ins.value
}

// equivalence check
func (ins *ProductId) Equals(value *ProductId) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}

func NewProductId(value string) (*ProductId, *errs.DomainError) {
	const LENGTH int = 36
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"

	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("productID must be %d characters long", LENGTH))
	}

	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("productID must be UUID format")
	}

	return &ProductId{value: value}, nil
}
