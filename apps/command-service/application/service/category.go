package service

import (
	"context"

	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
)

type CategoryService interface {
	Add(ctx context.Context, category *categories.Category) error
	Update(ctx context.Context, category *categories.Category) error
	Delete(ctx context.Context, category *categories.Category) error
}
