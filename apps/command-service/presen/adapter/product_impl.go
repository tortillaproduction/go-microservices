package adapter

import (
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/errs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductAdapterImpl struct{}

func NewProductAdapterImpl() ProductAdapter {
	return &ProductAdapterImpl{}
}

func (ins *ProductAdapterImpl) ToEntity(param *pbv1.ProductUpParam) (*products.Product, error) {
	switch param.GetCrud() {
	case pbv1.CRUD_INSERT:
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		id, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		product, err := products.NewProduct(name, price, categories.BuildCategory(id, nil))
		if err != nil {
			return nil, err
		}
		return product, nil

	case pbv1.CRUD_UPDATE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		cid, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, name, price, categories.BuildCategory(cid, nil)), nil

	case pbv1.CRUD_DELETE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, nil, nil, nil), nil

	default:
		return nil, errs.NewInternalError("recieved an unknown CRUD operation")
	}
}

func (ins *ProductAdapterImpl) ToResult(result any) *pbv1.ProductUpResult {
	var up_product *pbv1.Product
	var up_err *pbv1.Error

	switch v := result.(type) {
	case *products.Product:
		var c *pbv1.Category
		if v.Category() == nil {
			c = &pbv1.Category{Id: "", Name: ""}
		} else {
			c = &pbv1.Category{Id: v.Category().Id().Value(), Name: ""}
		}
		var name string = ""
		if v.Name() != nil {
			name = v.Name().Value()
		}
		var price int32 = 0
		if v.Price() != nil {
			price = int32(v.Price().Value())
		}
		up_product = &pbv1.Product{Id: v.Id().Value(), Name: name, Price: price, Category: c}

	case *errs.DomainError:
		up_err = &pbv1.Error{Type: "Domain Error", Message: v.Error()}

	case *errs.CRUDError:
		up_err = &pbv1.Error{Type: "CRUD Error", Message: v.Error()}

	case *errs.InternalError:
		up_err = &pbv1.Error{Type: "Internal Error", Message: "Internal Server Error"}
	}

	return &pbv1.ProductUpResult{
		Product:   up_product,
		Error:     up_err,
		Timestamp: timestamppb.Now(),
	}
}
