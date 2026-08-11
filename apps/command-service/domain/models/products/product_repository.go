package products

import (
	"context"
	"database/sql"
)

type ProductRepository interface {
	Exists(ctx context.Context, tran *sql.Tx, product *Product) error
	Create(ctx context.Context, tran *sql.Tx, product *Product) error
	UpdateById(ctx context.Context, tran *sql.Tx, product *Product) error
	DeleteById(ctx context.Context, tran *sql.Tx, product *Product) error
}
