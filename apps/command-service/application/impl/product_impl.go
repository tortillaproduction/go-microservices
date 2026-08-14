package impl

import (
	"context"

	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/products"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
)

type productServiceImpl struct {
	rep products.ProductRepository
	transaction
}

func NewProductServiceImpl(rep products.ProductRepository) service.ProductService {
	return &productServiceImpl{rep: rep}
}

func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.Exists(ctx, tran, product); err != nil {
		return err
	}

	if err = ins.rep.Create(ctx, tran, product); err != nil {
		return err
	}

	return nil
}

func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.UpdateById(ctx, tran, product); err != nil {
		return err
	}

	return err
}

func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.DeleteById(ctx, tran, product); err != nil {
		return err
	}

	return err
}
