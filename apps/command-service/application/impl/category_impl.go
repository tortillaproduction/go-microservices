package impl

import (
	"context"

	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/domain/models/categories"
)

type categoryServiceImpl struct {
	rep categories.CategoryRepository
	transaction
}

func NewCategoryServiceImpl(rep categories.CategoryRepository) service.CategoryService {
	return &categoryServiceImpl{rep: rep}
}

func (ins *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		defer ins.complete(tran, err)
	}()

	if err = ins.rep.Exists(ctx, tran, category); err != nil {
		return err
	}

	if err = ins.rep.Create(ctx, tran, category); err != nil {
		return err
	}

	return nil
}

func (ins *categoryServiceImpl) Update(ctx context.Context, category *categories.Category) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.UpdateById(ctx, tran, category); err != nil {
		return err
	}

	return err
}

func (ins *categoryServiceImpl) Delete(ctx context.Context, category *categories.Category) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.DeleteById(ctx, tran, category); err != nil {
		return err
	}

	return err
}
