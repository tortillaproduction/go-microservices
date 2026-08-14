package adapter

import (
	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
)

type CategoryAdapter interface {
	ToEntity(param *pbv1.CategoryUpParam) (*categories.Category, error)
	ToResult(result any) *pbv1.CategoryUpResult
}
