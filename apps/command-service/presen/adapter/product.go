package adapter

import (
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
)

type ProductAdapter interface {
	ToEntity(param *pbv1.ProductUpParam) (*products.Product, error)
	ToResult(result any) *pbv1.ProductUpResult
}
