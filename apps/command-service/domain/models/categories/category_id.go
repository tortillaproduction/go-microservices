package categories

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// value object (UUID)
type CategoryId struct {
	value string
}

func (ins *CategoryId) Value() string {
	return ins.value
}

// equivalence check
func (ins *CategoryId) Equals(value *CategoryId) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}

func NewCategoryId(value string) (*CategoryId, *errs.DomainError) {
	const LENGTH int = 36
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"

	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("categoryID must be %d characters long", LENGTH))
	}

	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("categoryID must be UUID format")
	}

	return &CategoryId{value: value}, nil
}
