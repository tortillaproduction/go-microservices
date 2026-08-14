package adapter

import (
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryAdapterImpl struct{}

func NewCategoryAdapterImpl() CategoryAdapter {
	return &CategoryAdapterImpl{}
}

func (ins *CategoryAdapterImpl) ToEntity(param *pbv1.CategoryUpParam) (*categories.Category, error) {
	switch param.GetCrud() {
	case pbv1.CRUD_INSERT:
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		category, err := categories.NewCategory(name)
		if err != nil {
			return nil, err
		}
		return category, nil

	case pbv1.CRUD_UPDATE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, name), nil

	case pbv1.CRUD_DELETE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, nil), nil

	default:
		return nil, errs.NewDomainError("recieved an unknown CRUD operation")
	}
}

func (ins *CategoryAdapterImpl) ToResult(result any) *pbv1.CategoryUpResult {
	var up_category *pbv1.Category
	var up_err *pbv1.Error

	switch v := result.(type) {
	case *categories.Category:
		if v.Name() == nil {
			up_category = &pbv1.Category{Id: v.Id().Value(), Name: ""}
		} else {
			up_category = &pbv1.Category{Id: v.Id().Value(), Name: v.Name().Value()}
		}

	case *errs.DomainError:
		up_err = &pbv1.Error{Type: "Domain Error", Message: v.Error()}

	case *errs.CRUDError:
		up_err = &pbv1.Error{Type: "CRUD Error", Message: v.Error()}

	case *errs.InternalError:
		up_err = &pbv1.Error{Type: "Internal Error", Message: "Internal Server Error"}
	}

	return &pbv1.CategoryUpResult{
		Category:  up_category,
		Error:     up_err,
		Timestamp: timestamppb.Now(),
	}
}
