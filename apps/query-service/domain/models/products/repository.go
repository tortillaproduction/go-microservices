package products

import "context"

type ProductRepository interface {
	List(ctx context.Context) ([]*Product, error)
	FindByProductId(ctx context.Context, productId string) (*Product, error)
	FinbByProductName(ctx context.Context, productName string) ([]*Product, error)
}
