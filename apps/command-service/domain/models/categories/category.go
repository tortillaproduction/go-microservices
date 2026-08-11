package categories

import (
	"github.com/google/uuid"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
)

// entity
type Category struct {
	id   *CategoryId
	name *CategoryName
}

func (ins *Category) Id() *CategoryId {
	return ins.id
}

func (ins *Category) Name() *CategoryName {
	return ins.name
}

// change value
func (ins *Category) ChangeCategoryName(name *CategoryName) {
	ins.name = name
}

// equivalence check
func (ins *Category) Equals(obj *Category) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("nil specified as an argument")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

func NewCategory(name *CategoryName) (*Category, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewCategoryId(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &Category{
				id:   id,
				name: name,
			}, nil
		}
	}
}

func BuildCategory(id *CategoryId, name *CategoryName) *Category {
	return &Category{
		id:   id,
		name: name,
	}
}
