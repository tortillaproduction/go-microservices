package categories

import "context"

type CategoryRepository interface {
	List(ctx context.Context) ([]*Category, error)
	FindByCategoryId(ctx context.Context, categoryId string) (*Category, error)
}
