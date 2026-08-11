package categories

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Exists(ctx context.Context, tran *sql.Tx, category *Category) error
	Create(ctx context.Context, tran *sql.Tx, category *Category) error
	UpdateById(ctx context.Context, tran *sql.Tx, category *Category) error
	DeleteById(ctx context.Context, tran *sql.Tx, category *Category) error
}
